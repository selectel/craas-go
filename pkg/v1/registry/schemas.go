package registry

import (
	"encoding/json"
	"time"
)

// Registry represents an unmarshalled registry from an API response.
type Registry struct {
	// ID is a unique identifier of the registry.
	ID string `json:"id"`

	// Name is a name of the registry.
	Name string `json:"name"`

	// CreatedAt is a timestamp in UTC timezone of when the registry has been created.
	CreatedAt time.Time `json:"createdAt"`

	// Status is a status of the registry.
	Status Status `json:"status"`

	// Size is a registry storage usage in bytes.
	Size int64 `json:"size"`

	// SizeLimit is a registry storage limit in bytes.
	SizeLimit int64 `json:"sizeLimit"`

	// Used is a registry storage percentage usage.
	Used float32 `json:"used"`
}

// Status represents a custom type for various registry statuses.
type Status string

const (
	StatusActive   Status = "ACTIVE"
	StatusCreating Status = "CREATING"
	StatusDeleting Status = "DELETING"
	StatusGC       Status = "GARBAGE_COLLECTION"
	StatusError    Status = "ERROR"
	StatusUnknown  Status = "UNKNOWN"
)

func getSupportedStatuses() []Status {
	return []Status{
		StatusActive,
		StatusCreating,
		StatusDeleting,
		StatusGC,
		StatusError,
	}
}

func isStatusSupported(s Status) bool {
	for _, v := range getSupportedStatuses() {
		if s == v {
			return true
		}
	}

	return false
}

func (result *Registry) UnmarshalJSON(b []byte) error {
	type tmp Registry
	var s struct {
		tmp
	}

	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	*result = Registry(s.tmp)

	// Check cluster status.
	if !isStatusSupported(s.tmp.Status) {
		result.Status = StatusUnknown
	}

	return nil
}
