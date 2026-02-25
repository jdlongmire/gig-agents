package auth

// IsOperator returns true if the given GitHub username is in the operator allowlist.
func IsOperator(username string, allowedOperators []string) bool {
	for _, op := range allowedOperators {
		if op == username {
			return true
		}
	}
	return false
}

// DetermineRole returns "operator" if the username is in the allowlist,
// otherwise "client". Once persisted to the DB, role never changes via this function.
func DetermineRole(username string, allowedOperators []string) string {
	if IsOperator(username, allowedOperators) {
		return "operator"
	}
	return "client"
}
