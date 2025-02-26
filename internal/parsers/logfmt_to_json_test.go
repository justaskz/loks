package parsers

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLogfmtToJson(t *testing.T) {
	Convey("", t, func() {
		line := "level=info message=\"Package sent\" ip=127.0.0.1"
		expected := map[string]string{
			"level":   "info",
			"message": "Package sent",
			"ip":      "127.0.0.1",
		}
		expectedJson, _ := json.Marshal(expected)

		log := LogfmtToJson(line)
		So(log, ShouldEqual, string(expectedJson))
	})
}
