/*
Given a string of digits,
you should replace any digit below 5 with '0'
and any digit 5 and above with '1'. Return the resulting string.

Note: input will never be an empty string
*/

package kata

func FakeBin(x string) string {
	var newString string
	for _, chr := range x {
		if int(chr)-'0' < 5 {
			newString += "0"
		} else {
			newString += "1"
		}
	}
	return newString
}
