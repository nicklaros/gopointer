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
			value := float32(22.22)
			result := Float32Pointer(value)

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

	Convey("Test Int16Pointer", t, func() {
		Convey("When called", func() {
			value := int16(99)
			result := Int16Pointer(value)

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
}
