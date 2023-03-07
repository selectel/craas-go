package gc

// StartGCOpts represents options for starting a garbage collection.
type StartGCOpts struct {
	// DeleteUntagged is a flag that indicates whether to delete untagged images.
	DeleteUntagged bool
}
