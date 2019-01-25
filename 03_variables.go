//Variables and variable definition. Static, dynamic and mixed declaration. Variable error declarion. Defining multiple variables. Constants.

package main

import "fmt"

func main() {
	// syntax: var variableName variableType
	// names are always written in camelCase or CamelCase (even constants)
	// lowercase first letter means that variable can't be used outside the package
	// var packagePrivate string
	// uppercase - variable can be imported by other packages
	// var CanBeImported string

	// declaring a variable - memory is allocated and no value is specified by the programmer, default value is assigned
	var declared int
	fmt.Printf("Default value for integers is: %d\n", declared)

	// defining a variable - memory is allocated and a value, specified by the programmer, is assigned
	var hello string = "Hello World!"
	fmt.Println(hello)

	// go can infer types based on the value on the right hand side of assignemt
	helloInferred := "Value on the right hand side of assignemt is of type string"
	fmt.Println(helloInferred)

	// assigning a new value to
	helloInferred = "When we know that var has been defined, we use '=' for all following assignemts"
	fmt.Println(helloInferred)

	// further assignments should use '='
	// uncomment the following line and run the program
	// helloInferred := "helloInferred already exists"

	// static type system - types don't typically change during program execution
	// uncomment the following line and run the program
	// helloInferred = 42

	// defining multiple vars
	// it is possible to declare multiple variables in a block
	// var (
	// 	variable1 string
	// 	variable2 string
	// )

	// constants
	const pi = 3.14159

	// constant block
	const (
		constant1 = 1
		constant2 = "constant"
	)
}
