package testing

import "github.com/selectel/craas-go/pkg/v1/token"

const testTokenID = `CRgAAAAAWiMnwN63eyASwQk8a3DBPRPirT9fWQTx`

// testCreateTokenResponseRaw represents a raw token create response.
const testCreateTokenResponseRaw = `{
    "expireAt": 1666644533,
    "expireIn": 43200,
    "token": "CRgAAAAAWiMnwN63eyASwQk8a3DBPRPirT9fWQTx"
}`

var expectedCreateTokenResponse = &token.Token{
	Token:     testTokenID,
	ExpiresAt: 1666644533,
	ExpiresIn: 43200,
}

const testGetTokenResponseRaw = `{
    "expireAt": 1666644533,
    "expireIn": 43200
}`

var expectedGetTokenResponse = &token.Token{
	ExpiresAt: 1666644533,
	ExpiresIn: 43200,
	Token:     testTokenID,
}

// testRefreshTokenResponseRaw represents a raw token refresh response.
const testRefreshTokenResponseRaw = `{
      "expireAt": 1666649999,
      "expireIn": 43200,
      "token": "CRgAAAAAWiMnwN63eyASwQk8a3DBPRPirT9fWQTx"
}`

// expectedRefreshTokenResponse represents an expected token refresh response.
var expectedRefreshTokenResponse = &token.Token{
	Token:     testTokenID,
	ExpiresAt: 1666649999,
	ExpiresIn: 43200,
}
