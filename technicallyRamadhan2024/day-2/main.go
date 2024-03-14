package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(CountBits(1234))
}

func CountBits(n uint) int {
	count := 0
	//generate code to turn number to binary
	//count the number of 1s in the binary
	//return the count
	binary := strconv.FormatInt(int64(n), 2)
	for _, v := range binary {
		if v == '1' {
			count++
		}
	}

	return count
}
