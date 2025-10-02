package main

import "fmt"

func main() {
	a := "scissors"
	b := "paper"
	c := "rock"

	fmt.Println(a > b)
	fmt.Println(b > c)
	fmt.Println(a > c)
}
