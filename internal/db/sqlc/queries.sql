-- name: GetAccountByGitHubID :one
SELECT * FROM accounts WHERE github_id = ? LIMIT 1;

-- name: GetAccountByID :one
SELECT * FROM accounts WHERE id = ? LIMIT 1;

-- name: GetAccountByUsername :one
SELECT * FROM accounts WHERE github_username = ? LIMIT 1;

-- name: CreateAccount :one
INSERT INTO accounts (id, github_id, github_username, display_name, avatar_url, role, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, datetime('now'), datetime('now'))
RETURNING *;

-- name: UpdateAccount :one
UPDATE accounts
SET github_username = ?,
    display_name    = ?,
    avatar_url      = ?,
    updated_at      = datetime('now')
WHERE id = ?
RETURNING *;

-- name: CreateSession :one
INSERT INTO sessions (id, account_id, token_hash, expires_at, created_at)
VALUES (?, ?, ?, ?, datetime('now'))
RETURNING *;

-- name: GetSessionByTokenHash :one
SELECT * FROM sessions WHERE token_hash = ? AND expires_at > datetime('now') LIMIT 1;

-- name: DeleteSession :exec
DELETE FROM sessions WHERE id = ?;

-- name: DeleteExpiredSessions :exec
DELETE FROM sessions WHERE expires_at <= datetime('now');

-- name: DeleteSessionsByAccountID :exec
DELETE FROM sessions WHERE account_id = ?;
