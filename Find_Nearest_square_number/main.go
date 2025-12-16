/*
Your task is to find the nearest square number of a positive integer n. In mathematics,
a square number or perfect square is an integer that is the square of an integer; in other words,
it is the product of some integer with itself.

For example, if n = 111, then the nearest square number equals 121, since 111 is closer to 121, the square of 11,
than 100, the square of 10.

If n is already a perfect square (e.g. n = 144, n = 81, etc.), you need to just return n.

Good luck :)
*/

package main

import (
	"fmt"
	"math"
)

func NearestSq(n int) int {
	sqrt := math.Sqrt(float64(n))
	if n == int(sqrt)*int(sqrt) {
		return n
	} else {
		lower := int(math.Floor(sqrt))
		upper := lower + 1

		lowerSquare := lower * lower
		upperSquare := upper * upper

		diffLower := n - lowerSquare
		diffUpper := upperSquare - n

		if diffLower <= diffUpper {
			return lowerSquare
		}
		return upperSquare
	}
}

func main() {
	fmt.Println(NearestSq(2))
}
