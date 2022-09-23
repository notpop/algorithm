package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func StrStdin() (stringInput string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringInput = scanner.Text()

	stringInput = strings.TrimSpace(stringInput)
	return
}

func IntStdin() (int, error) {
	stringInput := StrStdin()

	return strconv.Atoi(stringInput)
}

func main() {
	i, err := IntStdin()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i * i)
	}
}
