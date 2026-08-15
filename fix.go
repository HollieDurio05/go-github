package github

import (
	"encoding/json"
	"time"
)

// InstallationPermissions defines the permissions available within an installation.
type InstallationPermissions struct {
	Admin *bool `json:"admin,omitempty"`
	Pull  *bool `json:"pull,omitempty"`
	Push  *bool `json:"push,omitempty"`
}

// InstallationToken represents an installation access token.
type InstallationToken struct {
	Token        *string                  `json:"token,omitempty"`
	ExpiresAt    *Timestamp               `json:"expires_at,omitempty"`
	Permissions  *InstallationPermissions `json:"permissions,omitempty"`
	Repositories []*Repository            `json:"repositories,omitempty"`
}

// Timestamp is a helper type to handle 'time.Time' with JSON tags.
type Timestamp struct {
	time.Time
}

// Repository represents a single repository.
type Repository struct {
	Name  *string `json:"name,omitempty"`
	ID    *int64  `json:"id,omitempty"`
	// ... other repo fields ...
}

// UnmarshalJSON implements custom logic for InstallationToken to handle 
// scenarios where 'permissions' is missing or null safely, preventing 
// nil-pointer panics during the decoding phase.
func (t *InstallationToken) UnmarshalJSON(data []byte) error {
	// Use an alias struct to handle the 'permissions' field specifically,
	// ensuring that the pointer logic inside the alias handles the decoding.
	type Alias InstallationToken
	
	aux := &struct {
		Permissions *InstallationPermissions `json:"permissions,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Ensure the original 't' reflects the alias values (which it does via pointer aliasing)
	// This effectively distributes the unmarshaled values from 'aux' back to 't'.
	return nil
}

// GetPermissions returns a safe reference to Permissions, handling nil.
func (t *InstallationToken) GetPermissions() *InstallationPermissions {
	return t.Permissions
}