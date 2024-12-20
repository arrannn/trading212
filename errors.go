package trading212

import "fmt"

// ValidationError represents a specific trading validation error
type ValidationError struct {
	Code    string `json:"code"`
	Message string `json:"clarification"`
}

// Error implements the error interface
func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}
