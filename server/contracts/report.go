package contracts

// Report represents the result of the operation.
type Report struct {
	Success bool    `json:"success"`
	Errors  []Error `json:"errors,omitempty"`
}

func (r *Report) AddError(code string, field string, message string) {

	newError := Error{Code: code, Field: field, Message: message}

	r.Errors = append(r.Errors, newError)
}
