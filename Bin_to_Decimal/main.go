/*
Complete the function which converts a binary number
(given as a string) to a decimal number.
*/

package main

import (
	"fmt"
	"strconv"
)

func BinToDec(bin string) int {
	conv, _ := strconv.ParseInt(bin, 2, 64)
	return int(conv)
}

func main() {
	fmt.Println(BinToDec("101010"))
}
