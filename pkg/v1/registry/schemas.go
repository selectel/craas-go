package registry

import "time"

// Registry represents an unmarshalled registry from an API response.
type Registry struct {
	// ID is a unique identifier of the registry.
	ID string `json:"id"`

	// Name is a name of the registry.
	Name string `json:"name"`

	// CreatedAt is a timestamp in UTC timezone of when the registry has been created.
	CreatedAt time.Time `json:"createdAt"`

	// Status is a status of the registry.
	Status string `json:"status"`

	// Size is a registry storage usage in bytes.
	Size int64 `json:"size"`

	// SizeLimit is a registry storage limit in bytes.
	SizeLimit int64 `json:"sizeLimit"`

	// Used is a registry storage percentage usage.
	Used float32 `json:"used"`
}
