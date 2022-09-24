package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	result := "No"
	for index := a; index <= b; index++ {
		if 100%index == 0 {
			result = "Yes"
		}
	}

	fmt.Println(result)
}
