package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/justaskz/loks/internal/parsers"
)

type JSONData map[string]interface{}

func isJSON(line string) bool {
	var js JSONData
	return json.Unmarshal([]byte(line), &js) == nil
}
func printJson(json string) {
	json = DetectLogLevel(json)
	fmt.Println(json)
}

func isLogFmt(line string) bool {
	return strings.Contains(line, "=") && !isJSON(line)
}

func print(line string) {
	fmt.Println(line)
}

func DetectLogLevel(jsonStr string) string {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		return jsonStr
	}

	logLevels := map[string]string{
		"error": "\033[31merror\033[0m", // Red
		"warn":  "\033[33mwarn\033[0m",  // Orange/Yellow
		"info":  "\033[32minfo\033[0m",  // Green
		"debug": "\033[34mdebug\033[0m", // Blue
	}

	if level, exists := data["level"].(string); exists {
		if coloredLevel, ok := logLevels[strings.ToLower(level)]; ok {
			data["level"] = coloredLevel
		}
	}

	modifiedJSON, err := json.Marshal(data)
	if err != nil {
		return jsonStr
	}

	return string(modifiedJSON)
}

func printLogFmt(line string) {
	json := parsers.LogfmtToJson(line)
	printJson(json)
}

func chooseLogger(line string) func(string) {
	switch {
	case isJSON(line):
		return printJson
	case isLogFmt(line):
		return printLogFmt
	default:
		return print
	}
}

func main() {
	var scanner *bufio.Scanner

	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	var firstLine string
	if scanner.Scan() {
		firstLine = scanner.Text()
	}

	logger := chooseLogger(firstLine)
	logger(firstLine)
	for scanner.Scan() {
		log := scanner.Text()
		logger(log)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}
