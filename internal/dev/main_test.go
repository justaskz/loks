package dev

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {
	Convey("", t, func() {
		So(true, ShouldEqual, true)
	})
}
