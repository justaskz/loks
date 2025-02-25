package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type JSONData map[string]interface{}

func isJSON(s string) bool {
	var js JSONData
	return json.Unmarshal([]byte(s), &js) == nil
}

func isLogFmt(s string) bool {
	return strings.Contains(s, "=") && !isJSON(s)
}

func logFmtToJson(logFmtStr string) string {
	re := regexp.MustCompile(`(\w+)=([^\s]+)`)
	matches := re.FindAllStringSubmatch(logFmtStr, -1)

	data := make(map[string]string)
	for _, match := range matches {
		if len(match) == 3 {
			data[match[1]] = match[2]
		}
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "{}"
	}

	return string(jsonBytes)
}

func printLog(str string) {
	fmt.Println(str)
}

func printLogFmt(str string) {
	log := logFmtToJson(str)
	fmt.Println(log)
}

// func chooseLogger(line string) func() {
// 	return func() {}
// }

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

	switch {
	case isJSON(firstLine):
		printLog(firstLine)
		for scanner.Scan() {
			log := scanner.Text()
			printLog(log)
		}
	case isLogFmt(firstLine):
		printLogFmt(firstLine)

		for scanner.Scan() {
			log := scanner.Text()
			printLogFmt(log)
		}
	default:
		fmt.Println(firstLine)

		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}
