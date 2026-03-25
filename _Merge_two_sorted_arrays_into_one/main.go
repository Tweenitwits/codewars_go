/*
You are given two sorted arrays that contain only integers.
These arrays may be sorted in either ascending or descending order.
Your task is to merge them into a single array, ensuring that:

The resulting array is sorted in ascending order.

Any duplicate values are removed, so each integer appears only once.

If both input arrays are empty, return an empty array.

No input validation is needed, as both arrays are guaranteed to contain zero or more integers.

Examples (input -> output)
* [1, 2, 3, 4, 5], [6, 7, 8, 9, 10] -> [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

* [1, 3, 5, 7, 9], [10, 8, 6, 4, 2] -> [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

* [1, 3, 5, 7, 9, 11, 12], [1, 2, 3, 4, 5, 10, 12] -> [1, 2, 3, 4, 5, 7, 9, 10, 11, 12]
Happy coding!
*/

package main

import (
	"fmt"
	"sort"
)

funt MergeArrays(arr1, arr2 []int) []int {
	result := []int{}
	uniqueMap := make(map[int]struct{})
	combined := append(arr1, arr2...)

	for _, val := range combined {
		if _, exists := uniqueMap[val]; !exists {
			uniqueMap[val] = struct{}{}
			result = append(result, val)
		}
	}

	sort.Ints(result)

	return result
}

// func MergeArrays(arr1, arr2 []int) []int {
// 	lenArrs := len(arr1) + len(arr2)
// 	uniqueMap := make(map[int]struct{}, lenArrs)
// 	result := make([]int, 0, lenArrs)

// 	for _, val := range arr1 {
// 		if _, exists := uniqueMap[val]; !exists {
// 			uniqueMap[val] = struct{}{}
// 			result = append(result, val)
// 		}
// 	}

// 	for _, val := range arr2 {
// 		if _, exists := uniqueMap[val]; !exists {
// 			uniqueMap[val] = struct{}{}
// 			result = append(result, val)
// 		}
// 	}
// 	sort.Ints(result)

// 	return result

// }

func main() {
	arr1 := []int{1, 3, 5, 7, 9, 11, 12}
	arr2 := []int{1, 2, 3, 4, 5, 10, 12}
	fmt.Print(MergeArrays(arr1, arr2))
}
