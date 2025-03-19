package unwrap

import (
	"errors"
)

// UnwrapRaw simply returns the value or panics if there's an error.
func UnwrapRaw[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

// Unwrap checks for nil values, but only for pointer types.
func Unwrap[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}

	// This avoids nil checking for non-pointer types
	if isNil(value) {
		panic(errors.New("nil value returned"))
	}
	return value
}

// isNil checks if a generic value is nil, but only for pointer or interface types.
func isNil[T any](value T) bool {
	// Convert to interface{} before checking for nil
	return (any(value) == nil)
}
