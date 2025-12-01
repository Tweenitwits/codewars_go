/*
Complete the function that takes two integers (a, b, where a < b)
and return an array of all integers between the input parameters,
including them.

For example:

a = 1
b = 4
--> [1, 2, 3, 4]
*/

package main

import "fmt"

func Between(a, b int) []int {
	numSlice := []int{}
	for i := a; i <= b; i++ {
		numSlice = append(numSlice, i)
	}
	return numSlice
}

func main() {
	fmt.Println(Between(1, 4))
}
