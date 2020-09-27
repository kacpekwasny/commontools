package commontools

import (
	"fmt"
)

// SameType checks if args are of the same type
func SameType(values ...interface{}) bool {
	// Handle no items
	if len(values) < 2 {
		return true
	}

	var (
		areSameType bool   = true
		prevStr     string = fmt.Sprintf("%T", values[0])
		nowStr      string
	)

	// Check if all Types are the same
	for _, value := range values[1:] {
		nowStr = fmt.Sprintf("%T", value)
		if nowStr != prevStr {
			return false
		}
		prevStr = nowStr
	}

	return areSameType
}

// Type returns string of the type of first arg
func Type(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
