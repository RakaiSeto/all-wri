package main

import (
	"fmt"
	"strings"
)

func main() {
	spinWords("Hey fellow warriors")
	spinWords("This is a test")
	spinWords("This is another test")
}

func spinWords(input string) {
	output := ""

	arrayWords := strings.Split(input, " ")

	for _, word := range arrayWords {
		if len(word) > 4 {
			innerOutput := ""
			arrayChar := strings.Split(word, "")

			for i := len(word); i > 0; i-- {
				innerOutput += arrayChar[i-1]
			}
			output += innerOutput
		} else {
			output += word
		}
		output += " "
	}
	output = strings.TrimSuffix(output, " ")
	fmt.Println(output)
}
