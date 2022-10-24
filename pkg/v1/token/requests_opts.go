package token

import "fmt"

type TTL string

const (
	TTL1Year   TTL = "1y"
	TTL12Hours TTL = "12h"
)

var ErrInvalidTokenTTL = fmt.Errorf("invalid token ttl")

// CreateOpts represents options for the token create request.
type CreateOpts struct {
	// TokenTTL is a token expiration duration.
	TokenTTL TTL
}
