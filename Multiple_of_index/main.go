/*
Return a new array consisting of elements which are multiple of their own index in input array (length > 1).

Some cases:
[22, -6, 32, 82, 9, 25] =>  [-6, 32, 25]

[68, -1, 1, -7, 10, 10] => [-1, 10]

[-56,-85,72,-26,-14,76,-27,72,35,-21,-67,87,0,21,59,27,-92,68] => [-85, 72, 0, 68]
*/

package main

import "fmt"

func multipleOfIndex(ints []int) []int {
	idxSlice := []int{}
	for i := 1; i <= len(ints)-1; i++ {
		if ints[i]%i == 0 {
			idxSlice = append(idxSlice, ints[i])
		}
	}
	return idxSlice
}

func main() {
	fmt.Println(multipleOfIndex([]int{28, 38, -44, -99, -13, -54, 77, -51}))
}
