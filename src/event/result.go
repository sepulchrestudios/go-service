package event

// Result represents the outcome of processing an event.
type Result struct {
	// Success indicates whether the event was processed successfully.
	success bool

	// Return holds any relevant data returned from processing the event.
	returnData any

	// Error is an error (if any) encountered during event processing.
	err error
}

// Error retrieves the error message encountered during event processing. This also allows the Result struct to be
// used as an error type.
func (r *Result) Error() string {
	if r == nil || r.err == nil {
		return ""
	}
	return r.err.Error()
}

// ErrorInstance retrieves any error encountered during event processing.
func (r *Result) ErrorInstance() error {
	if r == nil {
		return nil
	}
	return r.err
}

// Return retrieves any relevant data returned from processing the event.
func (r *Result) Return() any {
	if r == nil {
		return nil
	}
	return r.returnData
}

// Success indicates whether the event was processed successfully.
func (r *Result) Success() bool {
	if r == nil {
		return false
	}
	return r.success
}

// NewEmptyResult creates a new empty Result instance.
func NewEmptyResult() *Result {
	return &Result{}
}

// NewResult creates a new Result instance.
func NewResult(success bool, returnData any, err error) *Result {
	return &Result{
		success:    success,
		returnData: returnData,
		err:        err,
	}
}
