//Structures

package main

import "fmt"

//A struct is a collection of fields.
//Each field has a name and a type.
//Create a struct with 3 fields
type Pony struct {
	Name  string
	Color string
	Age   int
}

func (p Pony) Rename(newName string) {
	p.Name = newName
}

func main() {

	//Create a value of that type.
	var twilightSparkle Pony
	twilightSparkle.Name = "Twilight Sparkle"
	twilightSparkle.Color = "Purple"
	twilightSparkle.Age = 3

	fmt.Println(twilightSparkle)

	//Struct fields are accessed using a dot.
	fmt.Println(twilightSparkle.Name)
	twilightSparkle.Color = "White"
	fmt.Println(twilightSparkle)

	twilightSparkle.Rename("Spike")
	fmt.Println(twilightSparkle.Name)
	//Q: Why does it print the original Name? How to fix it?
}
