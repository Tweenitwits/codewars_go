/*
The code provided is supposed replace all the dots . in the specified String str with dashes -

But it's not working properly.

Task
Fix the bug so we can all go home early.

Notes
String str will never be null.
*/

package main

import (
	"fmt"
	"regexp"
)

// replaces all "." with "-"
func ReplaceDots(str string) string {
	return regexp.MustCompile(`\.`).ReplaceAllString(str, "-")
}

func main() {
	fmt.Println(ReplaceDots("one.two.three"))
}
