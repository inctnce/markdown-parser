// go run main.go

package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func scanTextFile(path string) (string, error) {

	data, err := ioutil.ReadFile("test.md")
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func getWords(text string) (string, int) {
	words := []byte{}
	i := 0
	for true {
		if text[i] == '\n' {
			break
		}
		words = append(words, text[i])
		i++
	}
	return string(words), i
}

func parseDefinitions(text string) {
	length := len(text)
	keys := []string{}
	values := []string{}

	for i := 0; i < length-3; i++ {
		if text[i:i+3] == "###" {
			i += 4

			key, inc1 := getWords(text[i:])
			keys = append(keys, key)
			i += inc1 + 2

			value, inc2 := getWords(text[i:])
			values = append(values, value)
			i += inc2 + 2
		}
	}
	for i := 0; i < len(keys); i++ {
		fmt.Print(keys[i], "\n", values[i], "\n\n")
	}
}

func main() {
	text, err := scanTextFile("test.md")
	if err != nil {
		log.Fatal(err)
	}
	parseDefinitions(text)
}
