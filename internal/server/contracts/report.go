package contracts

// Report represents the result of the operation.
type Report struct {
	Success bool    `json:"success"`
	Errors  []Error `json:"errors,omitempty"`
}

func (r *Report) AddError(code int, field string, parameters ...interface{}) {

	var newError Error

	if len(parameters) > 0 {
		newError = CreateError(code, field, parameters...)
	} else {
		newError = CreateError(code, field)
	}

	r.Errors = append(r.Errors, newError)
}
