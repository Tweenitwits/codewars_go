/*
Build a function that returns an array of integers from n to 1 where n>0.

Example : n=5 --> [5,4,3,2,1]
*/

package kata

func ReverseSeq(n int) []int {
	totalSlc := []int{}
	for i := n; i >= 1; i-- {
		totalSlc = append(totalSlc, i)
	}
	return totalSlc
}
