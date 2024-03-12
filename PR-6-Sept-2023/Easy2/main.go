package main

import (
	"fmt"
	"strings"
)

func main() {
	XO("ooxx")
	XO("xooxx")
	XO("ooxXm")
	XO("zpzpzpp")
	XO("zzoo")
}

func XO(input string) {
	var Xs []string
	var Os []string

	for _, s := range strings.ToLower(input) {
		if string(s) == "x" {
			Xs = append(Xs, string(s))
		} else if string(s) == "o" {
			Os = append(Os, string(s))
		}
	}

	if len(Xs) == len(Os) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
