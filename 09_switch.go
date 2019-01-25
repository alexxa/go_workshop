//Desicion making. Branching with switch.

package main

import "fmt"

func main() {

	// More complex decision making
	// Shorter verion of if - else
	// Cases evaluated top to bottom, first matching case is selected
	// No automatic fall through, instead keyword 'fallthrough' can be used

	// Switching on an expression
	switch "match" {
	case "match":
		fmt.Println("Matched!")
	case "noMatch":
		fmt.Println("This branch will not execute")
	default:
		// If no above cases match, default branch is used
		fmt.Println("Defaulted!")
	}

	// Switching with no condition, the same as switch true
	// Instead 'case' branches are evaluated based on 'truthfulness', top to bottom
	// Idiomatic way of writing if-else-if-else chains
	var myFavouriteFruit = "banana"
	switch {
	case myFavouriteFruit == "ketchup":
		fmt.Println("Blahhhhh!")
	case myFavouriteFruit == "banana":
		fmt.Println("They're so easy to peel!")
	default:
		// If no above cases match, default branch is used
		fmt.Println("Defaulted!")
	}

	//Task: Check the current time, and accordingly print out
	//Good morning, Good afternoon or Good evening.
	//Use time package and Now().

	// Falling through!
	switch {
	case true:
		fmt.Println("This one mathces!")
		fallthrough
	case true:
		fmt.Println("Even this one mathces!")
		fallthrough
	case false:
		fmt.Println("Yes, even this is selected.")
	}

	//Cases can be presented in comma-separated lists.
	fmt.Println(escape('?'))
}

func escape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+':
		return true
	}
	return false
}
