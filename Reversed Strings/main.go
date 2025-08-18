/*
Complete the solution so that it reverses the string passed into it.

'world'  =>  'dlrow'
'word'   =>  'drow'

*/

package kata

func Solution(word string) string {
	convRune := []rune(word)
	for i, j := 0, len(convRune)-1; i < j; i, j = i+1, j-1 {
		convRune[i], convRune[j] = convRune[j], convRune[i]
	}
	return string(convRune)
}
