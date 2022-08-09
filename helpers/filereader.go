package helpers

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func ReadFileText(path string) string {
	text, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(text)
}

func ReadFileLines(path string) []string {
	text := ReadFileText(path)
	text = strings.ReplaceAll(text, "\r\n", "\n")
	return strings.Split(text, "\n")
}

func ReadFileLinesAsInt(path string) []int {
	lines := ReadFileLines(path)
	result := make([]int, len(lines))
	for i, l := range lines {
		num, err := strconv.Atoi(l)
		if err != nil {
			log.Fatal(err)
		}
		result[i] = num
	}
	return result
}
