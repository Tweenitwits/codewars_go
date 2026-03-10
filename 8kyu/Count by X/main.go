/*
Create a function with two arguments
that will return an array of the first n multiples of x.

Assume both the given number
and the number of times to count will be positive numbers greater than 0.

Return the results as an array or list ( depending on language ).

Examples
x = 1, n = 10 --> [1,2,3,4,5,6,7,8,9,10]
x = 2, n = 5  --> [2,4,6,8,10]
*/

package kata

func CountBy(x, n int) []int {
	total_slc := []int{}
	for i := 1; i <= n*x; i++ {
		if i%x == 0 {
			total_slc = append(total_slc, i)
		}
	}
	return total_slc
}
