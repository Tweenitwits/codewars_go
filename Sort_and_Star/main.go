/*
You will be given a list of strings.
You must sort it alphabetically (case-sensitive,
and based on the ASCII values of the chars)
and then return the first value.

The returned value must be a string,
and have "***" between each of its letters.

You should not remove or add elements from/to the array.
*/

package main

import (
	"fmt"
	"sort"
)

func TwoSort(arr []string) string {
	total := ""
	sort.Strings(arr)
	for _, chr := range arr[0] {
		total += string(chr) + "***"
	}
	return total[:len(total)-3]
}

func main() {
	fmt.Println(TwoSort([]string{"bitcoin", "take", "over", "the", "world", "maybe", "who", "knows", "perhaps"}))
}
