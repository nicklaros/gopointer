package gopointer

import (
	"fmt"
	"time"
)

// BoolPointer returns pointer of provided boolean value.
func BoolPointer(value bool) *bool {
	return &value
}

// Float32Pointer returns pointer of provided float32 value.
func Float32Pointer(value float32) *float32 {
	return &value
}

// Float64Pointer returns pointer of provided float64 value.
func Float64Pointer(value float64) *float64 {
	return &value
}

// IntPointer returns pointer of provided int value.
func IntPointer(value int) *int {
	return &value
}

// Int8Pointer returns pointer of provided int8 value.
func Int8Pointer(value int8) *int8 {
	return &value
}

// Int16Pointer returns pointer of provided int16 value.
func Int16Pointer(value int16) *int16 {
	return &value
}

// Int32Pointer returns pointer of provided int32 value.
func Int32Pointer(value int32) *int32 {
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

// UintPointer returns pointer of provided uint value.
func UintPointer(value uint) *uint {
	return &value
}

// Uint8Pointer returns pointer of provided uint8 value.
func Uint8Pointer(value uint8) *uint8 {
	return &value
}

// Uint16Pointer returns pointer of provided uint16 value.
func Uint16Pointer(value uint16) *uint16 {
	return &value
}

// Uint32Pointer returns pointer of provided uint32 value.
func Uint32Pointer(value uint32) *uint32 {
	return &value
}

// Uint64Pointer returns pointer of provided uint64 value.
func Uint64Pointer(value uint64) *uint64 {
	gimmeYourAddress := &99

	fmt.Print(gimmeYourAddress)
	return &value
}
