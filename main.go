package main

import (
	"fmt"
	"get-to-work-problem/calculator"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func main() {
	input := readFileAsString("input.txt")
	params, err := calculator.ReadInputAsParam(input)
	if err != nil {
		panic(err)
	}

	for i, p := range params {
		result := p.CanEveryoneMakeItToWork()
		switch v := result.(type) {
		case string:
			fmt.Printf("Case # %d: %s\n", i+1, v)
		default:
			fmt.Printf("Case # %d: %v\n", i+1, strings.Trim(fmt.Sprint(v), "[]"))
		}
	}
}

func readFileAsString(filePath string) string {
	abs, err := filepath.Abs(filePath)
	if err != nil {
		panic(fmt.Sprintf("Error when setting file path - %s", err.Error()))
	}
	data, err := ioutil.ReadFile(abs)
	if err != nil {
		panic(fmt.Sprintf("Error when reading file - %s", err.Error()))
	}
	return string(data)
}
