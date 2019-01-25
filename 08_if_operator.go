//Desicion making. Logical if operator.

package main

import "fmt"

func main() {
	// In programming languages conditions are used to control flow of the program
	// Syntax: if condition { executed only if condition is truthful }
	// Optionally with 'else' branch
	// if condition { executed only if condition is truthful } else { executed only if condition is NOT truthful }
	if true {
		fmt.Println("Truth is Truth. Always.")
	} else {
		fmt.Println("Truth isn't Truth, The Ministry of Truth said.")
	}

	// variables valid only within the scope of the execution blocks
	if colour, err := isFavouriteColour(""); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(colour)
	}

	// variables valid within the scope of the enclosing function
	colour, err := isFavouriteColour("orange")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Is orange Josef's favourite colour: %t\n", colour)
}

func isFavouriteColour(colour string) (bool, error) {
	if colour == "" {
		return false, fmt.Errorf("Please send us a colour :(")
	}
	if colour == "orange" {
		return true, nil
	}
	return false, nil
}

//Task: Find if a number is odd or even and print out the result.
