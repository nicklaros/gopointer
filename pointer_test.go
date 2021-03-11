package gopointer

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPointerUtil(t *testing.T) {
	Convey("Test BoolPointer", t, func() {
		Convey("When called", func() {
			value := true
			result := BoolPointer(value)

			Convey("it should returns the right pointer", func() {
				So(result, ShouldResemble, &value)
			})
		})
	})

	Convey("Test Float32Pointer", t, func() {
		Convey("When called", func() {
			value := float32(99.99)
			result := Float32Pointer(value)

			Convey("it should returns the right pointer", func() {
				So(result, ShouldResemble, &value)
			})
		})
	})

	Convey("Test Float64Pointer", t, func() {
		Convey("When called", func() {
			value := float64(99.99)
			result := Float64Pointer(value)

			Convey("it should returns the right pointer", func() {
				So(result, ShouldResemble, &value)
			})
		})
	})

	Convey("Test IntPointer", t, func() {
		Convey("When called", func() {
			value := int(99)
			result := IntPointer(value)

			Convey("it should returns the right pointer", func() {
				So(result, ShouldResemble, &value)
			})
		})
	})

	Convey("Test Int8Pointer", t, func() {
		Convey("When called", func() {
			value := int8(99)
			result := Int8Pointer(value)

			Convey("it should returns the right pointer", func() {
				So(result, ShouldResemble, &value)
			})
		})
	})

	Convey("Test Int16Pointer", t, func() {
		Convey("When called", func() {
			value := int16(99)
			result := Int16Pointer(value)

			Convey("it should returns the right pointer", func() {
				So(result, ShouldResemble, &value)
			})
		})
	})

	Convey("Test Int32Pointer", t, func() {
		Convey("When called", func() {
			value := int32(99)
			result := Int32Pointer(value)

			Convey("it should returns the right pointer", func() {
				So(result, ShouldResemble, &value)
			})
		})
	})

	Convey("Test Int64Pointer", t, func() {
		Convey("When called", func() {
			value := int64(99)
			result := Int64Pointer(value)

			Convey("it should returns the right pointer", func() {
				So(result, ShouldResemble, &value)
			})
		})
	})

	Convey("Test StringPointer", t, func() {
		Convey("When called", func() {
			value := "Dummy String"
			result := StringPointer(value)

			Convey("it should returns the right pointer", func() {
				So(result, ShouldResemble, &value)
			})
		})
	})

	Convey("Test TimePointer", t, func() {
		Convey("When called", func() {
			value := time.Now()
			result := TimePointer(value)

			Convey("it should returns the right pointer", func() {
				So(result, ShouldResemble, &value)
			})
		})
	})

	Convey("Test UintPointer", t, func() {
		Convey("When called", func() {
			value := uint(99)
			result := UintPointer(value)

			Convey("it should returns the right pointer", func() {
				So(result, ShouldResemble, &value)
			})
		})
	})

	Convey("Test Uint8Pointer", t, func() {
		Convey("When called", func() {
			value := uint8(99)
			result := Uint8Pointer(value)

			Convey("it should returns the right pointer", func() {
				So(result, ShouldResemble, &value)
			})
		})
	})

	Convey("Test Uint16Pointer", t, func() {
		Convey("When called", func() {
			value := uint16(99)
			result := Uint16Pointer(value)

			Convey("it should returns the right pointer", func() {
				So(result, ShouldResemble, &value)
			})
		})
	})

	Convey("Test Uint32Pointer", t, func() {
		Convey("When called", func() {
			value := uint32(99)
			result := Uint32Pointer(value)

			Convey("it should returns the right pointer", func() {
				So(result, ShouldResemble, &value)
			})
		})
	})

	Convey("Test Uint64Pointer", t, func() {
		Convey("When called", func() {
			value := uint64(99)
			result := Uint64Pointer(value)

			Convey("it should returns the right pointer", func() {
				So(result, ShouldResemble, &value)
			})
		})
	})
}
