// Package constant defines constant values used throughout the application.
//
// It includes the definition of the `AppMode` type and constants for different
// application modes such as production, development, and test. Additionally,
// this package defines constants for common HTTP header fields and content types.

package constant

// AppMode represents the mode in which the application is running.
// It can be one of "production", "development", or "test".
type AppMode = string

// Constants for different application modes
const (
	production  AppMode = "production"  // Represents the production environment
	development AppMode = "development" // Represents the development environment
	test        AppMode = "test"        // Represents the test environment
)

// app holds the values for different application modes, enabling easy access
// to these constants across the application.
type app struct {
	Production  AppMode
	Development AppMode
	Test        AppMode
}

// App is a globally accessible variable that holds the different application modes.
var App = app{
	Production:  production,
	Development: development,
	Test:        test,
}

// ContentType is the HTTP header key used to specify the content type of a request or response.
const ContentType = "Content-Type"

// ApplicationJSON is the HTTP content type for JSON-formatted data.
const ApplicationJSON = "application/json"
