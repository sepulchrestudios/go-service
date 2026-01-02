package mail

// Result represents the outcome of processing a mail message.
type Result struct {
	// err is an error (if any) encountered during mail message processing.
	err error

	// returnData holds any relevant data returned from processing the mail message.
	returnData any

	// source is the source mail message associated with this result.
	source MessageContract

	// success indicates whether the mail message was processed successfully.
	success bool
}

// Error retrieves the error message encountered during mail message processing. This also allows the Result struct to
// be used as an error type.
func (r *Result) Error() string {
	if r == nil || r.err == nil {
		return ""
	}
	return r.err.Error()
}

// ErrorInstance retrieves any error encountered during mail message processing.
func (r *Result) ErrorInstance() error {
	if r == nil {
		return nil
	}
	return r.err
}

// Return retrieves any relevant data returned from processing the mail message.
func (r *Result) Return() any {
	if r == nil {
		return nil
	}
	return r.returnData
}

// Source returns the source mail message associated with this result.
func (r *Result) Source() MessageContract {
	if r == nil {
		return nil
	}
	return r.source
}

// Success indicates whether the mail message was processed successfully.
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
func NewResult(success bool, returnData any, err error, source MessageContract) *Result {
	return &Result{
		err:        err,
		returnData: returnData,
		source:     source,
		success:    success,
	}
}
