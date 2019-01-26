//Arrays. Declaring and accessing arrays.

package main

import "fmt"

func main() {
	//The type [n]T is an array of n values of a single type T.
	//An array has a fixed size, since its length is part of its type.
	var veggies [5]string
	veggies[0] = "Cucumber"
	veggies[1] = "Carrot"
	fmt.Println("Veggies list", veggies, "length", len(veggies))

	//Task: Create a simple loop with the for construct,
	//which fills an array [5]int with random odd integers and
	//print that array to the screen.

	//Set the array entries when the array is declared
	fruits := [4]string{"apple", "banana", "orange", "kiwi"}
	fmt.Println(fruits, len(fruits))

	//Length on initialization
	pasta := [...]string{"ziti", "macaroni", "raviolli", "rigatoni", "rotini"}
	fmt.Println(pasta, len(pasta))

	//Change a value of an array element
	fruits[0] = "tomato"
	fmt.Println(fruits)

	//"Out of bounds" Error
	//fmt.Println(fruits[10])

	//Array types are always one-dimensional
	//but may be composed to form multi-dimensional types.
	var matrix [3][3]int
	fmt.Println(matrix)

	//Initialization
	m := [2][2]int{[2]int{1, 2}, [2]int{3, 4}}
	fmt.Println(m)
	//Simplified syntax
	a := [2][2]int{{4, 3}, {2, 1}}
	fmt.Println(a)
}
