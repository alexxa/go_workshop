//Language Overview. Ask what the audience knows about Go. Stress Features.
//Environemnt Setup. Help. Docs. GOPATH. How to run Go programs. Verifying installation with "Hello World!"
//On "Hello World" explain Go program structure and syntax: package declaration and import, functions, variables, statements and experessions, comments, tokens, line separators, identifiers, keywords, whitespaces.
//Data Types: boolean, numeric, string, derived. Converting between types
//Variables and variable definition. Static, dynamic and mixed declaration. Variable error declarion. Defining multiple variables. Constants.
//Println and Printf functions.
//Local and global variables.
//Operators: arithmetic, relational, logical, (no bitwise!), assignment
//Pointers vs Values
//Desicion making. Logical if operator. Branching with switch and select statements.
//Loops. For loop. Nested loop. Loop control (break, continue, goto). Infinite loop.
//Functions. Functions definition, calling, usage. Functions returning multiple values. Variadic Functions. Closure. Recursion. Defer, panic and recover
//Strings and operations with strings.
//Arrays. Declaring and accessing arrays.
//Slices
//Structures
//Code organization.
//Errors handling

package main

import "fmt"

func main() {
	//godoc fmt Println.
	fmt.Println("Hello World")

	//Variables. Types of variables. How to declare and define a variable
	var apples = 5
	var bananas int = 10
	var oranges float32 = 2.5
	//var apples, bananas, oranges int
	//apples = 5
	fmt.Println(apples, bananas, oranges)
	var str = "Do you like apples or bananas?"
	fmt.Println(str)
	var isFruit bool = true
	fmt.Println(isFruit)

	//Arithmetic Operators.
	//Addition, substraction, multiplication, division, modulo division
	var x = 2
	var y = 3
	var sum int
	sum = x + y
	fmt.Println("Sum is", sum)

	//Constants. Comments. An error if a var is not used.
	//const pi = 3.14

	//About Printf function
	var e float64 = 2.7182818284
	fmt.Printf("%f \n", e)
	fmt.Printf("%.2f \n", e)
	fmt.Printf("%T \n", e)

	fmt.Printf("%t \n", isFruit)

	//Strings. Operations with Strings
	var message string = "Hello"
	fmt.Println(len(message))
	fmt.Println(message + " John and Jane.\nHow are you?")

	//Logical Operators
	hour := 9
	isWeekend = true
	if (hour < 9) || (hour > 18) {
		fmt.Println("The office is closed")
	} else if (hour > 12) && (hour < 13) {
		fmt.Println("It's lunch time. The office is closed")
	} else {
		fmt.Println("The office is open")
	}
	//Task: Modify this example to check weekends

	//Switch Operator

}
