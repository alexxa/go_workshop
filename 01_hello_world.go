/*
Go program structure:
  - package declaration and import
  - functions
  - variables
  - statements and experessions
  - comments

Syntax:
  - tokens
  - line separators
  - identifiers
  - keywords
  - whitespaces
*/
//Every Go program is made up of packages.
//Usage of main package makes it an executable program.
package main

//Usage of packages.
//The package name is the same as the last element of the import path.
import (
	"fmt"
	"math"
)

//Multiple import statements look like:
//import "fmt"
//import "math"
//import "maht/rand"

func main() {
	//go doc fmt Println
	fmt.Println("Hello World!")
	//Exported names start with a capital.
	//Unexported names, lowercase, are not accessible from outsude the package
	fmt.Println(math.Pi)
}
