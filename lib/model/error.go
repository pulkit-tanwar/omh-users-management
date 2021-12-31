package model

// Errors structure
type Errors struct {
	Error []Error `json:"error"`
}

// Error structure
type Error struct {
	ErrorCode        int    `json:"errorCode"`
	ErrorDescription string `json:"errorDescription"`
}

// NewErrorStructure is the constructor function to ErrorStructure
func NewErrorStructure(errorCode int, errorDescription string) *Error {
	return &Error{
		ErrorCode:        errorCode,
		ErrorDescription: errorDescription,
	}
}
