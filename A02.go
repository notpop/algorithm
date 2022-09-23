package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	result := "No"
	var input int
	for index := 0; index < n; index++ {
		fmt.Scan(&input)
		if input == x {
			result = "Yes"
		}
	}

	fmt.Println(result)
}
