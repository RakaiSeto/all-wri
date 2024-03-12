package main

import (
	"fmt"
	"strings"
)

func main() {
	accum("abcd")
	accum("RqaEzty")
	accum("cwAt")
}

func accum(input string) {
	output := ""

	charArray := strings.Split(input, "")
	for key, value := range charArray {
		repeater := key + 1
		for i := 0; i < repeater; i++ {
			if i == 0 {
				output += strings.ToUpper(value)
			} else {
				output += strings.ToLower(value)
			}
		}
		output += "-"
	}
	output = strings.TrimSuffix(output, "-")
	fmt.Println(output)
}
