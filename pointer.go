package gopointer

import "time"

// BoolPointer returns pointer of provided boolean value.
func BoolPointer(value bool) *bool {
	return &value
}

// Float32Pointer returns pointer of provided float32 value.
func Float32Pointer(value float32) *float32 {
	return &value
}

// IntPointer returns pointer of provided int value.
func IntPointer(value int) *int {
	return &value
}

// Int16Pointer returns pointer of provided int16 value.
func Int16Pointer(value int16) *int16 {
	return &value
}

// Int64Pointer returns pointer of provided int64 value.
func Int64Pointer(value int64) *int64 {
	return &value
}

// StringPointer returns pointer of provided string value.
func StringPointer(value string) *string {
	return &value
}

// TimePointer returns pointer of provided time value.
func TimePointer(value time.Time) *time.Time {
	return &value
}
