/*
Write a function that checks if a given string
(case insensitive) is a palindrome.

A palindrome is a word, number, phrase,
or other sequence of symbols that reads the same backwards as forwards,
such as madam or racecar.
*/

package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(str string) bool {
	runes := []rune(strings.ToLower(str))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	reversedString := string(runes)
	return str == reversedString
}

func main() {
	fmt.Println(IsPalindrome("a"))

}
