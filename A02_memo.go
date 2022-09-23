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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	n := IntStdinA02(scanner)
	x := IntStdinA02(scanner)

	scanner2 := bufio.NewScanner(os.Stdin)
	scanner2.Split(bufio.ScanWords)
	result := "No\n"
	for index := 0; index < n; index++ {
		i := IntStdinA02(scanner2)
		if x == i {
			result = "Yes\n"
		}
	}

	fmt.Println(result)
}
