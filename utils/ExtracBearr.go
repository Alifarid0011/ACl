package utils

import (
	"errors"
	"strings"
)

var (
	ErrMissingAuthHeader = errors.New("authorization header is missing")
	ErrInvalidAuthFormat = errors.New("invalid authorization header format")
)

func ExtractBearerToken(header string) (string, error) {
	if header == "" {
		return "", ErrMissingAuthHeader
	}
	parts := strings.Fields(header)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", ErrInvalidAuthFormat
	}
	return parts[1], nil
}
