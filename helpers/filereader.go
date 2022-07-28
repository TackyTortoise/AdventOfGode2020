package helpers

import (
	"io/ioutil"
	"log"
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
	return strings.Split(text, "\r\n")
}
