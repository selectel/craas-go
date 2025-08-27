package tokenv2

import "time"

type TokenV2 struct {
	ID         string     `json:"id,omitempty"`
	Name       string     `json:"name,omitempty"`
	CreatedAt  *time.Time `json:"createdAt,omitempty"`
	Expiration Expiration `json:"expiration"`
	Scope      Scope      `json:"scope"`
	Status     string     `json:"status,omitempty"`
	Token      string     `json:"token,omitempty"`
	LastUsedAt *time.Time `json:"lastUsedAt,omitempty"`
}

type Scope struct {
	ModeRW        bool     `json:"modeRW"`
	AllRegistries bool     `json:"allRegistries"`
	RegistryIDs   []string `json:"registryIds,omitempty"`
}

type Expiration struct {
	IsSet     bool      `json:"isSet"`
	ExpiresAt time.Time `json:"expiresAt,omitempty"`
}

type TokensV2 struct {
	Tokens     []TokenV2 `json:"tokens"`
	TotalCount int64     `json:"totalCount"`
}
