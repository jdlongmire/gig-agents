package config

import (
	"fmt"
	"os"
	"strings"
)

// Config holds all application configuration sourced from environment variables.
type Config struct {
	Port               string
	DBPath             string
	JWTSecret          string
	GitHubClientID     string
	GitHubClientSecret string
	AllowedOperators   []string
	BaseURL            string
}

// Load reads configuration from environment variables and returns a validated Config.
// Required vars: CREWPORT_JWT_SECRET, CREWPORT_GITHUB_CLIENT_ID, CREWPORT_GITHUB_CLIENT_SECRET
func Load() (*Config, error) {
	cfg := &Config{
		Port:    getEnv("CREWPORT_PORT", "8080"),
		DBPath:  getEnv("CREWPORT_DB_PATH", "./crewport.db"),
		BaseURL: getEnv("CREWPORT_BASE_URL", "http://localhost:8080"),
	}

	// Required vars
	cfg.JWTSecret = os.Getenv("CREWPORT_JWT_SECRET")
	cfg.GitHubClientID = os.Getenv("CREWPORT_GITHUB_CLIENT_ID")
	cfg.GitHubClientSecret = os.Getenv("CREWPORT_GITHUB_CLIENT_SECRET")

	var missing []string
	if cfg.JWTSecret == "" {
		missing = append(missing, "CREWPORT_JWT_SECRET")
	}
	if cfg.GitHubClientID == "" {
		missing = append(missing, "CREWPORT_GITHUB_CLIENT_ID")
	}
	if cfg.GitHubClientSecret == "" {
		missing = append(missing, "CREWPORT_GITHUB_CLIENT_SECRET")
	}
	if len(missing) > 0 {
		return nil, fmt.Errorf("missing required environment variables: %s", strings.Join(missing, ", "))
	}

	// Parse allowed operators list
	rawOps := getEnv("CREWPORT_ALLOWED_OPERATORS", "bobbyhiddn,jdlongmire")
	for _, op := range strings.Split(rawOps, ",") {
		op = strings.TrimSpace(op)
		if op != "" {
			cfg.AllowedOperators = append(cfg.AllowedOperators, op)
		}
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
