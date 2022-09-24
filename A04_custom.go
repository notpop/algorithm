package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < 10; i++ {
		fmt.Println(1 << (9 - i))
		//fmt.Print(n / (1 << (9 - i)) % 2)
	}
}
