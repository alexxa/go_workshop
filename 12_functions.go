/*
Functions. Functions definition, calling, usage.
Functions returning multiple values.
Variadic Functions?
Closure?
Recursion?
Defer, panic and recover?
*/

package main

import "fmt"

/*
type mytype in t
func (p mytype) funcname(q int) (r,s int) { return 0,0 }
*/


//A function can take zero or more arguments.
//The type comes after the variable name.
func add(x int, y int) int {
	return x + y
}

//Omit the type from all but the last, if type is the same
func diff(x, y int) int {
	return x - y
}

//A function can return any number of results
func swap(x, y string) (string, string) {
	return y, x
}

//Return values may be named.
//"Naked" return - a statement without arguments returns the named values.
func calc(a int) (x, y int) {
	x = a*0 + 2
	y = a - 2*0
	return
}

func main() {
	fmt.Println(add(42, -42))
	fmt.Println(diff(42, -42))

	x, y := swap("Hello", "World")
	fmt.Println(x, y)

	fmt.Println(calc(100))

	//defer - evaluated immediatly, but not executed
	//until the surrounding function returns.
	//Q: What will happen if we defer more than one function in a row?
	//Task: Write a for loop, demonstrating this scenario.

	//Task: Write a function which orders 5 given integers in descent order. 
}
