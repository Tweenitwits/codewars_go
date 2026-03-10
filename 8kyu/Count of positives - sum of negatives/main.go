/*
Given an array of integers.

Return an array,
where the first element is the count of positives numbers
and the second element is sum of negative numbers.
0 is neither positive nor negative.

If the input is an empty array or is null, return an empty array.

Example
For input [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15],
you should return [10, -65].
*/

package kata

func CountPositivesSumNegatives(numbers []int) []int {
	var res []int
	var positives, negatives int
	for _, val := range numbers {
		if val < 0 {
			negatives += val
		} else if val > 0 {
			positives += 1
		}
	}
	res = append(res, positives)
	res = append(res, negatives)
	return res
}
