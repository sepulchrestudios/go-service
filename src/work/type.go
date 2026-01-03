package work

// WorkType represents the type or category of a work item.
type WorkType string

const (
	// WorkTypeAll represents all (or no specific) work type(s).
	//
	// This is useful for subscribing to all work items but may not be appropriate when publishing them.
	WorkTypeAll WorkType = "all"
)
