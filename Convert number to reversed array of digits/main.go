/*
Given a random non-negative number,
you have to return the digits of this number within an array
in reverse order.

Example (Input => Output):
35231 => [1,3,2,5,3]
0     => [0]
*/

package main

import (
	"fmt"
	"strconv"
)

func Digitize(n int) []int {
	str := strconv.Itoa(n)

	var res string
	for _, v := range str {
		res = string(v) + res
	}

	var digits []int
	for _, char := range res {
		digits = append(digits, int(char-'0'))
	}
	return digits
}

func main() {
	fmt.Println(Digitize(35231))
}
