package utils

import (
	"acl-casbin/config"
	"acl-casbin/validation"
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"regexp"
	"strconv"
	"strings"
)

// ValidationError represents a detailed structure for a validation error in a form
// that includes the field, message, and additional metadata like index in arrays or slices.
type ValidationError struct {
	// Property is the name of the field that failed validation.
	Property string `json:"property"`
	// Tag is the validation tag that was violated (e.g., "required").
	// It is optional and may not always be present.
	Tag string `json:"tag,omitempty"`
	// Field represents the struct field name that failed validation.
	// It helps in identifying which field has the error.
	Field string `json:"field,omitempty"`
	// Value is the value of the field that failed validation.
	Value interface{} `json:"value,omitempty"`
	// Index holds the index in case the field is part of an array or slice.
	// If the field is not part of an array or slice, it is nil.
	Index *int `json:"index,omitempty"`
	// Message is the validation error message that describes the error in human-readable terms.
	Message string `json:"message"`
}

// GetValidationErrors extracts and converts validation errors into a structured format.
// It maps the validation errors into a slice of ValidationError objects, including the field name, validation tag, error message, and index for array/slice fields.
func GetValidationErrors(err error) *[]ValidationError {
	var validationErrors []ValidationError
	// Only process if there is an error
	if err != nil {
		var ve validator.ValidationErrors

		// Check if the error is of type ValidationErrors from the validator package
		if errors.As(err, &ve) {
			for _, err := range ve {
				// Prepare the error details
				var el ValidationError
				el.Property = err.Field()                                   // The field name
				el.Field = err.StructField()                                // The field's struct name
				el.Message = err.Translate(validation.TranslatorInstance()) // Translated message
				el.Index = parseIndex(err.StructNamespace())                // Parse index if it's in an array/slice
				// Append the formatted error to the validation errors list
				validationErrors = append(validationErrors, el)
			}
			return &validationErrors
		}
	}
	return nil
}

// parseIndex extracts the index of a field if the field is part of an array or slice.
// The index is extracted from the struct namespace (e.g., "Items[0].Name" would give index 0).
// Returns a pointer to the index integer or nil if no index is found.
func parseIndex(namespace string) *int {
	openBracket := strings.LastIndex(namespace, "[")  // Find the opening bracket '['
	closeBracket := strings.LastIndex(namespace, "]") // Find the closing bracket ']'
	// Ensure the brackets are valid and in the correct order
	if openBracket != -1 && closeBracket != -1 && closeBracket > openBracket {
		// Extract the index string between the brackets
		indexStr := namespace[openBracket+1 : closeBracket]
		// Attempt to parse the index string into an integer
		if index, err := strconv.Atoi(indexStr); err == nil {
			return &index
		}
	}
	return nil // Return nil if no valid index is found
}

func RegisterCustomValidators(v *validator.Validate, client *mongo.Client) {
	DB := client.Database(config.Get.Mongo.DbName)
	users := DB.Collection("users")
	_ = v.RegisterValidation("iran_mobile", IranianMobileValidator)
	v.RegisterValidation("unique_username", UniqueFieldValidator(users, "username"))
	v.RegisterValidation("unique_email", UniqueFieldValidator(users, "email"))
	v.RegisterValidation("unique_mobile", UniqueFieldValidator(users, "mobile"))
	_ = v.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()
		var (
			hasMinLen  = len(password) >= 6
			hasUpper   = false
			hasLower   = false
			hasNumber  = false
			hasSpecial = false
		)

		for _, ch := range password {
			switch {
			case 'a' <= ch && ch <= 'z':
				hasLower = true
			case 'A' <= ch && ch <= 'Z':
				hasUpper = true
			case '0' <= ch && ch <= '9':
				hasNumber = true
			case strings.ContainsRune("!@#$%^&*()-_=+[]{}|;:',.<>?/`~\\\"", ch):
				hasSpecial = true
			}
		}

		return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
	})
}

func UniqueFieldValidator(collection *mongo.Collection, field string) validator.Func {
	return func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		filter := bson.M{field: value}
		count, err := collection.CountDocuments(context.TODO(), filter)
		if err != nil {
			return false
		}
		return count == 0
	}
}

func IranianMobileValidator(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	match, _ := regexp.MatchString(`^09\d{9}$`, mobile)
	return match
}
