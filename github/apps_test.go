package github

import (
	"encoding/json"
	"testing"
)

func TestInstallationToken_MissingPermissions(t *testing.T) {
	rawJSON := `{"token":"v1.1234567890demo","expires_at":"2023-01-01T00:00:00Z"}`

	var token InstallationToken
	err := json.Unmarshal([]byte(rawJSON), &token)
	if err != nil {
		t.Fatalf("Unmarshal returned unexpected error: %v", err)
	}

	if token.Permissions != nil {
		t.Errorf("Expected Permissions to be nil, got %+v", token.Permissions)
	}
}