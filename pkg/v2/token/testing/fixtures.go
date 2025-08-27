package testing

import (
	"time"

	tokenV2 "github.com/selectel/craas-go/pkg/v2/token"
)

const testTokenID = `CRgAAAAAWiMnwN63eyASwQk8a3DBPRPirT9fWQTx`

// testCreateTokenResponseRaw represents a raw token create response.
const testCreateTokenResponseRaw = `{
  "id": "c29e3f63-0711-4772-a415-ad79973bdaef",
  "name": "my-token",
  "createdAt": "2023-02-13T15:00:00Z",
  "expiration": {
    "isSet": true,
    "expiresAt": "2030-01-01T00:00:00Z"
  },
  "scope": {
    "modeRW": true,
    "allRegistries": false,
    "registryIds": [
      "888af692-c646-4b76-a234-81ca9b5bcafe",
      "6303699d-c2cd-40b1-8428-9dcd6cc3d00d"
    ]
  },
  "status": "active",
  "token": "CRgAAAAAWiMnwN63eyASwQk8a3DBPRPirT9fWQTx"
}`

const testPatchTokenResponseRaw = `{
  "id": "c29e3f63-0711-4772-a415-ad79973bdaef",
  "name": "token",
  "createdAt": "2023-02-13T15:00:00Z",
  "expiration": {
    "isSet": true,
    "expiresAt": "2030-01-01T00:00:00Z"
  },
  "scope": {
    "modeRW": true,
    "allRegistries": false,
    "registryIds": [
    	"888af692-c646-4b76-a234-81ca9b5bcafe",
		"6303699d-c2cd-40b1-8428-9dcd6cc3d00d"
    ]
  },
  "lastUsedAt": "2023-02-14T15:25:10Z",
  "status": "active"
}`

// testListTokensResponseRaw represents a raw list token response.
const testListTokensResponseRaw = `
{
  "tokens": [
  	{
      "id": "c29e3f63-0711-4772-a415-ad79973bdaef",
      "name": "my-token",
      "createdAt": "2023-02-13T15:00:00Z",
      "expiration": {
        "isSet": true,
        "expiresAt": "2030-01-01T00:00:00Z"
      },
      "scope": {
        "modeRW": true,
        "allRegistries": false,
        "registryIds": [
          "888af692-c646-4b76-a234-81ca9b5bcafe",
          "6303699d-c2cd-40b1-8428-9dcd6cc3d00d"
        ]
      },
      "lastUsedAt": "2023-02-14T15:25:10Z",
      "status": "active"
    },
	  	{
      "id": "c29e3f63-0711-4772-a415-ad79973bdaef",
      "name": "my-token",
      "createdAt": "2023-02-13T15:00:00Z",
      "expiration": {
        "isSet": true,
        "expiresAt": "2030-01-01T00:00:00Z"
      },
      "scope": {
        "modeRW": true,
        "allRegistries": false,
        "registryIds": [
          "888af692-c646-4b76-a234-81ca9b5bcafe",
          "6303699d-c2cd-40b1-8428-9dcd6cc3d00d"
        ]
      },
      "lastUsedAt": "2023-02-14T15:25:10Z",
      "status": "active"
    }
  ],
  "totalCount": 2
}
`

var createdAt, _ = time.Parse("2006-01-02T15:04:05Z", "2023-02-13T15:00:00Z")
var expiresAt, _ = time.Parse("2006-01-02T15:04:05Z", "2030-01-01T00:00:00Z")
var lastUsedAt, _ = time.Parse("2006-01-02T15:04:05Z", "2023-02-14T15:25:10Z")

var Exp = tokenV2.Expiration{
	IsSet:     true,
	ExpiresAt: expiresAt,
}

var Scope = tokenV2.Scope{
	ModeRW:        true,
	AllRegistries: false,
	RegistryIDs: []string{
		"888af692-c646-4b76-a234-81ca9b5bcafe",
		"6303699d-c2cd-40b1-8428-9dcd6cc3d00d",
	},
}

var expectedCreateTokenResponse = &tokenV2.TokenV2{
	Token:      testTokenID,
	Status:     "active",
	ID:         "c29e3f63-0711-4772-a415-ad79973bdaef",
	Name:       "my-token",
	CreatedAt:  &createdAt,
	Expiration: Exp,
	Scope:      Scope,
}

var expectedPatchTokenResponse = &tokenV2.TokenV2{
	ID:         "c29e3f63-0711-4772-a415-ad79973bdaef",
	Name:       "token",
	CreatedAt:  &createdAt,
	Expiration: Exp,
	Scope:      Scope,
	LastUsedAt: &lastUsedAt,
	Status:     "active",
}

var expectedListTokenResponse = &tokenV2.TokensV2{
	TotalCount: 2,
	Tokens: []tokenV2.TokenV2{
		{
			Status:     "active",
			LastUsedAt: &lastUsedAt,
			ID:         "c29e3f63-0711-4772-a415-ad79973bdaef",
			Name:       "my-token",
			CreatedAt:  &createdAt,
			Expiration: struct {
				IsSet     bool      `json:"isSet"`
				ExpiresAt time.Time `json:"expiresAt,omitempty"`
			}{
				IsSet:     true,
				ExpiresAt: expiresAt,
			},
			Scope: struct {
				ModeRW        bool     "json:\"modeRW\""
				AllRegistries bool     "json:\"allRegistries\""
				RegistryIDs   []string "json:\"registryIds,omitempty\""
			}{
				ModeRW:        true,
				AllRegistries: false,
				RegistryIDs: []string{
					"888af692-c646-4b76-a234-81ca9b5bcafe",
					"6303699d-c2cd-40b1-8428-9dcd6cc3d00d",
				},
			},
		},
		{
			Status:     "active",
			LastUsedAt: &lastUsedAt,
			ID:         "c29e3f63-0711-4772-a415-ad79973bdaef",
			Name:       "my-token",
			CreatedAt:  &createdAt,
			Expiration: struct {
				IsSet     bool      `json:"isSet"`
				ExpiresAt time.Time `json:"expiresAt,omitempty"`
			}{
				IsSet:     true,
				ExpiresAt: expiresAt,
			},
			Scope: struct {
				ModeRW        bool     "json:\"modeRW\""
				AllRegistries bool     "json:\"allRegistries\""
				RegistryIDs   []string "json:\"registryIds,omitempty\""
			}{
				ModeRW:        true,
				AllRegistries: false,
				RegistryIDs: []string{
					"888af692-c646-4b76-a234-81ca9b5bcafe",
					"6303699d-c2cd-40b1-8428-9dcd6cc3d00d",
				},
			},
		},
	},
}
