package bdd

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	Convey("Given 2 even numbers", t, func() {
		a := 3
		b := 4
		Convey("When and the two numbersd", func() {
			c := a + b
			Convey("Then the result is still even", func() {
				So(c%2, ShouldEqual, 0)
			})
		})
	})
}
