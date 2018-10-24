package contracts

// Report represents the result of the operation.
type Report struct {
	Success bool    `json:"success"`
	Errors  []error `json:"errors,omitempty"`
}

func (r *Report) AddError(code string, field string, message string) {

	newError := error{Code: code, Field: field, Message: message}

	r.Errors = append(r.Errors, newError)
}
