package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func IntStdinA02(scanner *bufio.Scanner) int {
	scanner.Scan()
	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
	}

	return num
}

// なぜこのアプローチではACにならないのかよくわからない。
// ちょっと理由がわからなすぎるので一旦別のアプローチでここはスルーする
// 一応問題に乗ってる入力、出力の例では問題なく動いているので何かしらatcoderの形式に沿わない部分があるのだとは思う。

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	n := IntStdinA02(scanner)
	x := IntStdinA02(scanner)

	scanner2 := bufio.NewScanner(os.Stdin)
	scanner2.Split(bufio.ScanWords)
	result := "No"
	for index := 0; index < n; index++ {
		i := IntStdinA02(scanner2)
		if x == i {
			result = "Yes"
		}
	}

	fmt.Println(result)
}
