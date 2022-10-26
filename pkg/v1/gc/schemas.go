package gc

// GarbageSize represents an unmarshalled garbage size from an API response.
type GarbageSize struct {
	// NonReferenced is a size of the layers non-referenced to any repository digests.
	NonReferenced int64 `json:"sizeNonReferenced"`

	// Untagged is a size of the layers of the images with no tags.
	Untagged int64 `json:"sizeUntagged"`

	// Summary is a size of the sum of Untagged and NonReferenced image layers.
	Summary int64 `json:"sizeSummary"`
}
