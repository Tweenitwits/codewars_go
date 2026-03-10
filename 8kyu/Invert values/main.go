/*
Given a set of numbers, return the additive inverse of each.
Each positive becomes negatives, and the negatives become positives.

[1, 2, 3, 4, 5] --> [-1, -2, -3, -4, -5]
[1, -2, 3, -4, 5] --> [-1, 2, -3, 4, -5]
[] --> []

You can assume that all values are integers.
Do not mutate the input array.
*/

package kata

func Invert(arr []int) []int {
	total := []int{}
	for _, val := range arr {
		total = append(total, val*-1)
	}
	return total
}
