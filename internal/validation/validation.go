package validation

const (
	// NullOrEmpty indicates a missing field or property.
	// No args are expected.
	NullOrEmpty = 100

	// MaxLength indicates the max length of a string.
	// A integer is expected as an argument.
	MaxLength = 200

	// MinAndMaxRange indicates the min and lax length allowed for a string.
	// Two integers, min and max, are expected respectively as arguments.
	MinAndMaxRange = 201

	// InvalidFormat indicates that the received parameter is not formatted as expected.
	// No args are expected.
	InvalidFormat = 202

	// ExpectedFormat indicates a format expected for a parameter.
	// A string describing the format is expected as an argument.
	ExpectedFormat = 203

	// AlreadyRegistered indicates that the requested parameter is already registered.
	// No args are expected.
	AlreadyRegistered = 300

	// AlreadyActivated indicates that the requested parameter is already activated.
	// No args are expected.
	AlreadyActivated = 301

	// NotFound indicates that the requested parameter was not found.
	// No args are expected.
	NotFound = 302

	// CouldNotBeParsed indicates that an element or inputed data could not be processed.
	// No args are expected.
	CouldNotBeParsed = 900

	// InternalServerError indicates that an operation as caused an error internally.
	// No args are expected.
	InternalServerError = 1001
)

// Errors represents a map between error codes and messages.
var Errors = map[int]string{

	NullOrEmpty: "%s cannot be empty or null.",

	MaxLength:      "%s must have a maximum of %d characters.",
	MinAndMaxRange: "%s must have between %d and %d characters.",
	InvalidFormat:  "The specified %s is not in a valid format.",
	ExpectedFormat: "Expected format for %s is '%s'.",

	AlreadyRegistered: "%s already registered.",
	AlreadyActivated:  "%s already activated.",
	NotFound:          "%s not found.",

	CouldNotBeParsed: "The %s could not be parsed.",

	InternalServerError: "Oops! An internal error has ocurred. Please, try again later.",
}
