/*
For every good kata idea there seem to be quite a few bad ones!

In this kata you need to check the provided array (x) for good ideas 'good' and bad ideas 'bad'.
If there are one or two good ideas, return 'Publish!',
if there are more than 2 return 'I smell a series!'. If there are no good ideas,
as is often the case, return 'Fail!'.
*/

package main

import "fmt"

func Well(x []string) string {
	goodIdeas := 0
	for _, val := range x {
		if val == "good" {
			goodIdeas++
		}
	}
	if goodIdeas == 1 || goodIdeas == 2 {
		return "Publish!"
	} else if goodIdeas > 2 {
		return "I smell a series!"
	}
	return "Fail!"
}

func main() {
	fmt.Println(Well([]string{"bad", "bad", "bad"}))
}
