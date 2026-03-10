/*
Given an array of integers, return a new array with each value doubled.

For example:

[1, 2, 3] --> [2, 4, 6]
*/

package kata

func Maps(x []int) []int {
	var total []int
	for _, val := range x {
		total = append(total, val*2)
	}
	return total
}
