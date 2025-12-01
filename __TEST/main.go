package main

import (
	"fmt"
	"strconv"
)

func deleteOdd(num uint) uint {
	var numStr, res string
	var number int
	numStr = strconv.Itoa(int(num))
	for _, chr := range numStr {
		if (int(chr)-48)%2 == 0 && int(chr) != 48 {
			res += string(chr)
		}
	}
	if len(res) == 0 {
		number = 100
	} else {
		number, _ = strconv.Atoi(res)
	}
	return uint(number)
}

func main() {
	fmt.Println(deleteOdd(1013))
}
