package main

import (
	"fmt"
	"strconv"
)

func main() {
	add("99999999999999", "99999999999999")
}

func add(i1 string, i2 string) {
	int1, err := strconv.Atoi(i1)
	int2, err2 := strconv.Atoi(i2)

	if err != nil {
		panic(err)
	}
	if err2 != nil {
		panic(err2)
	}

	ans := int1 + int2

	fmt.Println(strconv.Itoa(ans))
}
