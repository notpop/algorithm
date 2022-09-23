package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func IntStdinB01() int {
	scanner.Scan()
	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
	}

	return num
}

func main() {
	scanner.Split(bufio.ScanWords)
	a := IntStdinB01()
	b := IntStdinB01()
	fmt.Println(a + b)
}
