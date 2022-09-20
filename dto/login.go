package dto

type LoginRequest struct {
	Username  string   `form:"username" json:"username"`
	Password  string   `form:"password" json:"password"`
	GrantType string   `form:"grant_type" json:"grant_type"`
	Scope     []string `form:"scope[]" json:"scope"`
}
type LoginResponse struct {
	AccessToken string   `json:"access_token"`
	TokenType   string   `json:"token_type"`
	ExpiresIn   int64    `json:"expires_in"`
	Scope       []string `json:"scope"`
}
