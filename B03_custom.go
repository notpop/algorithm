package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	result := "No"
	goods := make([]int, n)
	for index := 0; index < n; index++ {
		fmt.Scan(&goods[index])
	}

	for outerIndex := 0; outerIndex < n; outerIndex++ {
		for betweenIndex := outerIndex + 1; betweenIndex < n; betweenIndex++ {
			for innerIndex := betweenIndex + 1; innerIndex < n; innerIndex++ {
				if 1000 == goods[outerIndex]+goods[betweenIndex]+goods[innerIndex] {
					result = "Yes"
				}
			}
		}
	}

	fmt.Println(result)
}
