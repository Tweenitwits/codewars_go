/*
Let's play! You have to return which player won! In case of a draw return Draw!.

Examples(Input1, Input2 --> Output):

"scissors", "paper" --> "Player 1 won!"
"scissors", "rock" --> "Player 2 won!"
"paper", "paper" --> "Draw!"
*/

package kata

func Rps(p1, p2 string) string {
	win1 := "Player 1 won!"
	win2 := "Player 2 won!"
	switch {
	case p1 == "scissors" && p2 == "paper":
		return win1
	case p1 == "scissors" && p2 == "rock":
		return win2
	case p1 == "paper" && p2 == "scissors":
		return win2
	case p1 == "paper" && p2 == "rock":
		return win1
	case p1 == "rock" && p2 == "scissors":
		return win1
	case p1 == "rock" && p2 == "paper":
		return win2
	default:
		return "Draw"
	}
}
