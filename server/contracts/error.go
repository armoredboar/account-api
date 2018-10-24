package contracts

// Error indeicates the field and reason of an processing error.
type error struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
}
