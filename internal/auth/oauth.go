package auth

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	dbsqlc "github.com/bobbyhiddn/CrewPort/internal/db/sqlc"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

const (
	oauthStateCookie = "crewport_oauth_state"
	stateCookieTTL   = 10 * time.Minute
)

// OAuthHandler handles the GitHub OAuth redirect and callback flows.
type OAuthHandler struct {
	oauthCfg         *oauth2.Config
	jwtSecret        string
	allowedOperators []string
	queries          *dbsqlc.Querier
}

// NewOAuthHandler constructs an OAuthHandler.
func NewOAuthHandler(clientID, clientSecret, baseURL, jwtSecret string, allowedOperators []string, db *sql.DB) *OAuthHandler {
	cfg := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"read:user"},
		Endpoint:     github.Endpoint,
		RedirectURL:  baseURL + "/auth/github/callback",
	}
	return &OAuthHandler{
		oauthCfg:         cfg,
		jwtSecret:        jwtSecret,
		allowedOperators: allowedOperators,
		queries:          dbsqlc.New(db),
	}
}

// HandleRedirect sends the browser to GitHub's OAuth authorization page with
// a random state token stored in an httpOnly cookie.
func (h *OAuthHandler) HandleRedirect(w http.ResponseWriter, r *http.Request) {
	state, err := generateState()
	if err != nil {
		slog.Error("failed to generate OAuth state", "err", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     oauthStateCookie,
		Value:    state,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   int(stateCookieTTL.Seconds()),
	})

	url := h.oauthCfg.AuthCodeURL(state, oauth2.AccessTypeOnline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// HandleCallback processes the GitHub callback, exchanges the code for a token,
// fetches the user profile, upserts the account, and sets a JWT cookie.
func (h *OAuthHandler) HandleCallback(w http.ResponseWriter, r *http.Request) {
	// Validate state
	stateCookie, err := r.Cookie(oauthStateCookie)
	if err != nil || r.URL.Query().Get("state") != stateCookie.Value {
		slog.Warn("OAuth state mismatch")
		http.Error(w, "invalid OAuth state", http.StatusBadRequest)
		return
	}

	// Clear state cookie
	http.SetCookie(w, &http.Cookie{
		Name:     oauthStateCookie,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	})

	// Exchange code for access token
	code := r.URL.Query().Get("code")
	oauthToken, err := h.oauthCfg.Exchange(r.Context(), code)
	if err != nil {
		slog.Error("OAuth code exchange failed", "err", err)
		http.Error(w, "OAuth exchange failed", http.StatusBadGateway)
		return
	}

	// Fetch GitHub user info
	ghUser, err := fetchGitHubUser(r.Context(), h.oauthCfg, oauthToken)
	if err != nil {
		slog.Error("failed to fetch GitHub user", "err", err)
		http.Error(w, "failed to fetch user info", http.StatusBadGateway)
		return
	}

	// Upsert account
	account, err := h.upsertAccount(r.Context(), ghUser)
	if err != nil {
		slog.Error("failed to upsert account", "err", err, "github_username", ghUser.Login)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	slog.Info("user logged in", "username", account.GithubUsername, "role", account.Role)

	// Issue JWT
	tokenStr, err := CreateToken(h.jwtSecret, account.ID, account.GithubID, account.GithubUsername, account.Role)
	if err != nil {
		slog.Error("failed to create JWT", "err", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	SetTokenCookie(w, tokenStr)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

// gitHubUser is a minimal representation of the GitHub user API response.
type gitHubUser struct {
	ID        int64  `json:"id"`
	Login     string `json:"login"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}

// fetchGitHubUser calls the GitHub API /user endpoint with the given OAuth token.
func fetchGitHubUser(ctx context.Context, cfg *oauth2.Config, token *oauth2.Token) (*gitHubUser, error) {
	client := cfg.Client(ctx, token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return nil, fmt.Errorf("github api request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("github api returned %d", resp.StatusCode)
	}

	var user gitHubUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("decode github user: %w", err)
	}

	if user.Login == "" {
		return nil, errors.New("empty GitHub login returned")
	}

	return &user, nil
}

// upsertAccount creates a new account or updates an existing one.
// Role is only set on first creation — it never changes after that.
func (h *OAuthHandler) upsertAccount(ctx context.Context, ghUser *gitHubUser) (*dbsqlc.Account, error) {
	existing, err := h.queries.GetAccountByGitHubID(ctx, ghUser.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("get account: %w", err)
	}

	displayName := ghUser.Name
	if displayName == "" {
		displayName = ghUser.Login
	}

	if existing != nil {
		// Update mutable fields, preserve role
		updated, err := h.queries.UpdateAccount(ctx, dbsqlc.UpdateAccountParams{
			ID:             existing.ID,
			GithubUsername: ghUser.Login,
			DisplayName:    displayName,
			AvatarUrl:      ghUser.AvatarURL,
		})
		if err != nil {
			return nil, fmt.Errorf("update account: %w", err)
		}
		return updated, nil
	}

	// First login — assign role from allowlist
	role := DetermineRole(ghUser.Login, h.allowedOperators)
	created, err := h.queries.CreateAccount(ctx, dbsqlc.CreateAccountParams{
		ID:             uuid.New().String(),
		GithubID:       ghUser.ID,
		GithubUsername: ghUser.Login,
		DisplayName:    displayName,
		AvatarUrl:      ghUser.AvatarURL,
		Role:           role,
	})
	if err != nil {
		return nil, fmt.Errorf("create account: %w", err)
	}
	slog.Info("new account created", "username", created.GithubUsername, "role", created.Role)
	return created, nil
}

// generateState creates a cryptographically random state string for CSRF protection.
func generateState() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
