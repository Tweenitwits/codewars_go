/*
Write a function to convert a name into initials. This kata strictly takes two words with one space in between them.

The output should be two capital letters with a dot separating them.

It should look like this:

Sam Harris => S.H

patrick feeney => P.F
*/

package kata

import (
	"strings"
)

func AbbrevName(name string) string {
	tmp := strings.Split(name, " ")
	newWord := strings.ToUpper(string(tmp[0][0])) + "." + strings.ToUpper(string(tmp[1][0]))
	return newWord
}
