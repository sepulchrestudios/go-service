package work

// HandlerFunc defines the function signature for work item handler functions.
type HandlerFunc func(workItem WorkContract) WorkResultContract
