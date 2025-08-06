package utils

import (
	"fmt"
	"time"
)

const layout = time.RFC3339

func ParseDate(value string, fieldName string) (time.Time, error) {
	parsed, err := time.Parse(layout, value)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid %s format, expected RFC3339", fieldName)
	}
	return parsed, nil
}

func ParseOptionalDate(value *string, fieldName string) (*time.Time, error) {
	if value == nil {
		return nil, nil
	}
	parsed, err := time.Parse(layout, *value)
	if err != nil {
		return nil, fmt.Errorf("invalid %s format, expected RFC3339", fieldName)
	}
	return &parsed, nil
}
