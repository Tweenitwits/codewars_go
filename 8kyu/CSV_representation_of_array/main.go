/*
Create a function that returns the CSV representation of a two-dimensional numeric array.

Example:

input:
   [[ 0, 1, 2, 3, 4 ],
    [ 10,11,12,13,14 ],
    [ 20,21,22,23,24 ],
    [ 30,31,32,33,34 ]]

output:
     '0,1,2,3,4\n'
    +'10,11,12,13,14\n'
    +'20,21,22,23,24\n'
    +'30,31,32,33,34'
Array's length > 2.

More details here: https://en.wikipedia.org/wiki/Comma-separated_values

Note: you shouldn't escape the \n, it should work as a new line.
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ToCsvText(array [][]int) string {
	addElem := "\\n"
	resultStr := ""
	lastIndex := len(array) - 1
	for i, row := range array {
		stringsRow := []string{}
		for _, val := range row {
			tmp := strconv.Itoa(val)
			stringsRow = append(stringsRow, tmp)
		}
		s := strings.Join(stringsRow, ",")
		if i != lastIndex {
			resultStr += s + addElem
		} else {
			resultStr += s
		}
	}
	return resultStr
}

func main() {
	fmt.Println(ToCsvText([][]int{{0, 1, 2, 3, 4},
		{10, 11, 12, 13, 14},
		{20, 21, 22, 23, 24},
		{30, 31, 32, 33, 34}}))
}
