//Pointers vs Values.

package main

import "fmt"

func main() {
	//A pointer holds the memory address of a value.
	//The type *T is a pointer to a T value. Its zero value is nil.
	//The & operator generates a pointer to its operand.
	//The * operator denotes the pointer's underlying value.

	var a *int
	b := 42
	a = &b

	fmt.Println(a, b) //Q: What will it print?
	fmt.Println(*a)   // Q: What will it print?
	//*a = *a / 0
	//fmt.Println(b) //Q: What will it print?

}
