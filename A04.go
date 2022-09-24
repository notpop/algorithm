package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	var result string
	count := 0
	for n != 0 {
		result = strconv.Itoa(n%2) + result
		n /= 2
		count++
	}

	if count < 10 {
		var padding string
		for index := 0; index < 10-count; index++ {
			padding += "0"
		}
		result = padding + result
	}

	fmt.Println(result)
}
