//Strings and operations with strings.
//Package "strings".

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Working with strings!")

	fmt.Println("Green" + "banana")

	s := "Big red apple" //Strings are immutable
	fmt.Println("The length of s is", len(s))
	fmt.Println("The first symbol of s is", string(s[0]))
	fmt.Println("The last symbol of s is", string(s[len(s)-1]))

	if s[len(s)-1] == 'e' {
		fmt.Println("The last sybmol is e!")
	}

	//Task: print string "s" symbol by symbol, each on a new line.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c\n", s[i])
	}

	fmt.Println(s[0:4]) // Doesn't print the 4th symbol
	fmt.Println(s[0:])
	fmt.Println(s[:4])
	//What will happen if we omit both indeces and keep a column only?
	fmt.Println(s[:])

	substr := s[4:8]
	fmt.Println(substr)
}
