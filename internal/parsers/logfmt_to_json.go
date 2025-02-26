package parsers

import (
	"encoding/json"
	"strings"

	"github.com/go-logfmt/logfmt"
)

func LogfmtToJson(line string) string {
	decoder := logfmt.NewDecoder(strings.NewReader(line))
	data := make(map[string]interface{})

	for decoder.ScanRecord() {
		for decoder.ScanKeyval() {
			key := string(decoder.Key())
			value := string(decoder.Value())
			data[key] = value
		}
	}

	if err := decoder.Err(); err != nil {
		return "{}"
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "{}"
	}

	return string(jsonBytes)
}
