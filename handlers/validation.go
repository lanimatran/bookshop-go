package handlers

import (
	"fmt"
)

func ValidateJsonLength(json map[string]interface{}, expectedLength int) error {
	if (len(json) != expectedLength) {
		return fmt.Errorf("Expected %d argument", expectedLength)
	}

	return nil
}

func ValidateNonEmptyString(argName string, arg any) error {
	str, ok := arg.(string)
	if (!ok || len(str) == 0) {
		return fmt.Errorf("Expected non-empty string for argument %s", argName)
	}

	return nil
}

func ValidatePositiveNumber(argName string, arg any) error {
	val, ok := arg.(float64)
	if (!ok || val <= 0) {
		return fmt.Errorf("Expected positive number for argument %s", argName)
	}

	return nil
}
