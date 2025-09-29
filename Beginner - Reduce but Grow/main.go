/*
Given a non-empty array of integers,
return the result of multiplying the values together in order.
Example:

[1, 2, 3, 4] => 1 * 2 * 3 * 4 = 24
*/

package kata

func Grow(arr []int) int {
	total := 1
	for _, val := range arr {
		total *= val
	}
	return total
}
