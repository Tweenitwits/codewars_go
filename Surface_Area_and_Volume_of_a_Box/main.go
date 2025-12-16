/*
Write a function that returns the total surface area and volume of a box.

The given input will be three positive non-zero integers: width, height, and depth.

The output will be language dependant, so please check sample tests for the corresponding data type,
(list, tuple, struct, query, etcetera).
*/

package main

import "fmt"

func GetSize(w, h, d int) [2]int {
	resultSlice := []int{}
	areaBox := 2 * (w*h + w*d + d*h)
	volumeBox := w * h * d
	resultSlice = append(resultSlice, volumeBox)
	resultSlice = append(resultSlice, areaBox)
	return [2]int(resultSlice)
}

func main() {
	fmt.Println(GetSize(4, 2, 6))
}
