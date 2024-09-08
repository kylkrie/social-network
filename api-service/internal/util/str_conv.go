package util

import (
	"strconv"
)

// Int64ToString converts an int64 to a string
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// NullableInt64ToString converts a *int64 to a *string
func NullableInt64ToString(i *int64) *string {
	if i == nil {
		return nil
	}
	s := strconv.FormatInt(*i, 10)
	return &s
}

// StringToInt64 converts a string to an int64
func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// NullableStringToInt64 converts a *string to a *int64
func NullableStringToInt64(s *string) (*int64, error) {
	if s == nil {
		return nil, nil
	}
	i, err := strconv.ParseInt(*s, 10, 64)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

// StringToInt64MustParse converts a string to an int64, panicking on error
func StringToInt64MustParse(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

// NullableStringToInt64MustParse converts a *string to a *int64, panicking on error
func NullableStringToInt64MustParse(s *string) *int64 {
	if s == nil {
		return nil
	}
	i, err := strconv.ParseInt(*s, 10, 64)
	if err != nil {
		panic(err)
	}
	return &i
}
