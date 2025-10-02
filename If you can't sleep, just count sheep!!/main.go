/*
If you can't sleep, just count sheeps!!

Task:
Given a non-negative integer, 3 for example,
return a string with a murmur: "1 sheep...2 sheep...3 sheep...".
Input will always be valid, i.e. no negative integers.
*/

package kata

import (
	"fmt"
	"strconv"
)

func countSheep(num int) string {
	totalString := ""
	str := " sheep..."
	for i := 1; i <= num; i++ {
		totalString += (strconv.Itoa(i) + str)
	}
	return totalString
}

func main() {
	fmt.Println(countSheep(3))
}
