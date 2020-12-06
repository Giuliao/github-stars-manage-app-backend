// Package github extract data structure from github page
// https://docs.github.com/en/free-pro-team@latest/developers/apps/authorizing-oauth-apps
package github

// Auth auth data
type Auth struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	RedirectURI  string `json:"redirect_url"`
	State        string `json:"state"`
}

// AuthResp response data
type AuthResp struct {
	AccessToken string `json:"access_token,omitempy"`
	Scope       string `json:"scope,omitempty"`
	TokenType   string `json:"token_type,omitempty"`
}

// User user data
type User struct {
	AvatarURL  string `json:"avatar_url,omitempty"`
	StarredURL string `json:"starrted_url,omitempty"`
}
