package main

import (
	"fmt"
	"strconv"
)

// 別の人の提出を確認するともっと早い方法がありそう。
func main() {
	var n string
	fmt.Scan(&n)

	summary := 0
	start := 0
	for index := 1; index <= len(n); index++ {
		target, _ := strconv.Atoi(n[start:index])
		summary += 1 << (len(n) - index) * target
		start = index
	}

	fmt.Println(summary)
}
