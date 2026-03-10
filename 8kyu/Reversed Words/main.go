/*
Complete the solution so
that it reverses all of the words within the string passed in.

Words are separated by exactly one space and there are no leading
or trailing spaces.

Example(Input --> Output):

"The greatest victory is that which requires no battle" -->
"battle no requires which that is victory greatest The"
*/

package kata

import "strings"

func ReverseWords(str string) string {
	splitStr := strings.Split(str, " ")
	for i, j := 0, len(splitStr)-1; i < j; i, j = i+1, j-1 {
		splitStr[i], splitStr[j] = splitStr[j], splitStr[i]
	}
	return strings.Join(splitStr, " ")
}
