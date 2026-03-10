/*
Write a function to split a string and convert it into an array of words.

Examples (Input ==> Output):
"Robin Singh" ==> ["Robin", "Singh"]

"I love arrays they are my favorite" ==>
["I", "love", "arrays", "they", "are", "my", "favorite"]
*/

package kata

import "strings"

func StringToArray(str string) []string {
	return strings.Split(str, " ")
}
