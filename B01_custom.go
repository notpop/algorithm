package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func IntStdinB01Custom(scannerCustom *bufio.Scanner) int {
	scannerCustom.Scan()
	num, err := strconv.Atoi(scannerCustom.Text())
	if err != nil {
		fmt.Println(err)
	}

	return num
}

func main() {
	scannerCustom := bufio.NewScanner(os.Stdin)
	// ここで型確認したところscannerCustomだと*bufio.Scannerで&scannerCustomだと&scannerCustomが返却された
	// 実際にNewScannerを確認すると&Scannerが返却されてることから関数IntStdinB01Customの引数として渡す際はscannerCustomのまま渡すことが可能
	//fmt.Println(reflect.TypeOf(&scannerCustom))
	scannerCustom.Split(bufio.ScanWords)
	a := IntStdinB01Custom(scannerCustom)
	b := IntStdinB01Custom(scannerCustom)
	fmt.Println(a + b)
}
