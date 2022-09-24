package main

import "fmt"

// 何パターンか試したけどこのアプローチがatcoderでは最速だった
func main() {
	var n int
	fmt.Scan(&n)

	result := "No"
	goods := make([]int, n)
	for index := 0; index < n; index++ {
		fmt.Scan(&goods[index])
	}

	for aIndex, a := range goods {
		for bIndex, b := range goods {
			if aIndex == bIndex {
				continue
			}
			for cIndex, c := range goods {
				if aIndex == bIndex || aIndex == cIndex || bIndex == cIndex {
					continue
				}
				if 1000 == a+b+c {
					result = "Yes"
				}
			}
		}
	}

	fmt.Println(result)
}
