/*
Write a method, that will get an integer array
as parameter and will process every number from this array.

Return a new array with processing every number of the input-array like this:

If the number has an integer square root, take this, otherwise square the number.

Example
[4,3,9,7,2,1] -> [2,9,3,49,4,1]
Notes
The input array will always contain only positive numbers, and will never be empty or null.
*/

package main

import (
	"fmt"
	"math"
)

func SquareOrSquareRoot(arr []int) []int {
	totalArr := []int{}
	for _, val := range arr {
		sqrtFloat := math.Sqrt(float64(val))
		sqrtInt := int(sqrtFloat)
		if sqrtInt*sqrtInt == val {
			totalArr = append(totalArr, sqrtInt)
		} else {
			totalArr = append(totalArr, val*val)
		}
	}
	return totalArr
}

func main() {
	fmt.Println(SquareOrSquareRoot([]int{4, 3, 9, 7, 2, 1}))
}
