package github

import (
	"encoding/json"
)

// InstallationToken represents a GitHub Apps installation access token.
type InstallationToken struct {
	Token        *string                  `json:"token,omitempty"`
	ExpiresAt    *Timestamp               `json:"expires_at,omitempty"`
	Permissions  *InstallationPermissions `json:"permissions,omitempty"`
	Repositories []*Repository            `json:"repositories,omitempty"`
}

// GetToken returns the Token field if it's non-nil, zero value otherwise.
func (i *InstallationToken) GetToken() string {
	if i == nil || i.Token == nil {
		return ""
	}
	return *i.Token
}

// GetExpiresAt returns the ExpiresAt field if it's non-nil, zero value otherwise.
func (i *InstallationToken) GetExpiresAt() Timestamp {
	if i == nil || i.ExpiresAt == nil {
		return Timestamp{}
	}
	return *i.ExpiresAt
}

// GetPermissions returns the Permissions field if it's non-nil, zero value otherwise.
func (i *InstallationToken) GetPermissions() *InstallationPermissions {
	if i == nil {
		return nil
	}
	return i.Permissions
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (i *InstallationToken) UnmarshalJSON(data []byte) error {
	type Alias InstallationToken
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(i),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}