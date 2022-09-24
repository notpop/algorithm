package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)

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
					fmt.Println("Yes")
					os.Exit(0)
				}
			}
		}
	}

	fmt.Println("No")
}
