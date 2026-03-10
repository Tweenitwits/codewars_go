/*
Create a combat function that takes the player's current health and the amount of damage received,
and returns the player's new health. Health can't be less than 0.
*/

package main

import (
	"fmt"
	"reflect"
)

func combat(health, damage float64) float64 {
	if health > damage {
		return health - damage
	}
	return 0
}

func main() {
	fmt.Print(reflect.TypeOf(combat(6.0, 1.0)))
}
