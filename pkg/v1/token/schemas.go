package token

// Token represents  an unmarshalled token body from an API response.
type Token struct {
	// Token is a token string.
	Token string `json:"token"`

	// ExpiresAt is a token expiration time.
	ExpiresAt int64 `json:"expireAt"`

	// ExpiresIn is a token expiration time in seconds.
	ExpiresIn int64 `json:"expireIn"`
}
