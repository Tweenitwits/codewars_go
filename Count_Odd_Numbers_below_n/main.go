/*
Given a number n, return the number of positive odd numbers below n,
EASY!

Examples (Input -> Output)
7  -> 3 (because odd numbers below 7 are [1, 3, 5])
15 -> 7 (because odd numbers below 15 are [1, 3, 5, 7, 9, 11, 13])
Expect large Inputs!
*/

package main

import "fmt"

func OddCount(n int) int {
	res := 0
	for i := 1; i < n; i++ {
		if i%2 != 0 {
			res += 1
		}
	}
	return res
}

func main() {
	fmt.Println(OddCount(15))
}
