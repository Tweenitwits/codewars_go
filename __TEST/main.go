package main

import "fmt"

func create(iterations int) []int {
	a := make([]int, 0)
	for i := 0; i < iterations; i++ {
		a = append(a, i)
	}
	return a
}

func main() {
	sliceFromLoop()
	sliceFromLiteral()

}

func sliceFromLoop() {
	fmt.Printf("** NOT working as expected: **\n\n")
	i := create(11)
	fmt.Println("initial slice: ", i)
	j := append(i, 100)
	g := append(i, 101)
	h := append(i, 102)
	fmt.Printf("i: %v\nj: %v\ng: %v\nh:%v\n", i, j, g, h)
}

func sliceFromLiteral() {
	fmt.Printf("\n\n** working as expected: **\n")
	i := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("initial slice: ", i)
	j := append(i, 100)
	g := append(i, 101)
	h := append(i, 102)
	fmt.Printf("i: %v\nj: %v\ng: %v\nh:%v\n", i, j, g, h)
}
