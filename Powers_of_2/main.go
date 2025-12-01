/*
Complete the function that takes a non-negative integer n as input,
and returns a list of all the powers of 2
with the exponent ranging from 0 to n ( inclusive ).
*/

package main

import (
	"fmt"
	"math"
)

func PowersOfTwo(n int) []uint64 {
	power := []uint64{}
	for i := 0; i <= n; i++ {
		power = append(power, uint64(math.Pow(2, float64(i))))
	}
	return power
}

func main() {
	fmt.Println(PowersOfTwo(4))
}
