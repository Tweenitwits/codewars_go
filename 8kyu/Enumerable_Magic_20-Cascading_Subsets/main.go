/*
Create a method each_cons that accepts a list and a number n,
and returns cascading subsets of the list of size n, like so:

each_cons([1,2,3,4], 2)
  #=> [[1,2], [2,3], [3,4]]

each_cons([1,2,3,4], 3)
  #=> [[1,2,3],[2,3,4]]

As you can see, the lists are cascading; ie, they overlap, but never out of order.
*/

package main

import "fmt"

func eachCons(arr []int, n int) [][]int {
	if len(arr) == 0 || len(arr) < n {
		return [][]int{}
	}

	result := make([][]int, 0, len(arr)-n+1)

	for i := 0; i <= len(arr)-n; i++ {
		tmp := make([]int, n)
		copy(tmp, arr[i:i+n])
		result = append(result, tmp)
	}

	return result
}

func main() {
	arr := []int{3, 5, 8, 13}
	fmt.Println(eachCons(arr, 3))
}
