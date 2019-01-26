/*
Variables and variable definition.
Static, dynamic and mixed declaration.
Variable error declarion.
Defining multiple variables.
Default values.
Types.
*/

package main

import "fmt"

func main() {

	//Variables declaration one by one
	var conf string
	var year int
	var location string

	//Variables assignment
	conf = "DevConf.CZ"
	year = 2019
	location = "Brno"

	println(conf, year, location)
	//return

	//Variables declaration and initialization, one by one
	var event string = "DevConf.CZ" //full
	var date = 2020                 //data type will be autodefined
	place := "Brno"                 //no var keyword and type

	println(event, date, place)
	//return

	//Multiple declaration and initialization, on the same line, different types
	var title, duration, speakers = "Ready, Steady, Go!", 110, "I.Gulina, J.Karasek"

	println(title, duration, speakers)
	//return

	//The var statement declares a list of variables with the type declared last.
	var (
		book            string
		publicationDate int
		author          string
	)

	book = "The Little Prince"
	publicationDate = 1943
	author = "Antoine de Saint-Exupery"

	println(book, publicationDate, author)
	//return

	//Similar
	var (
		album, band string
		songs       int
	)

	album, band = "A Night At The Opera", "Queen"
	songs = 12

	println(album, band, songs)
	//return

	//A var declaration can include initializers, one per variable.
	var (
		workshop string = "Ready, Steady, Go!"
		room     string = "C228"
		capacity int    = 70
	)

	println(workshop, room, capacity)
	//return

	//Data types can be omitted here too
	var (
		x1 = 10
		x2 = "string"
		x3 = true
	)

	println(x1, x2, x3)
	//return

	//Declared and not used variable -> error

	//Default values
	var defaultInt int
	var defaultFloat float32
	var defaultString string
	var defaultBoolean bool
	fmt.Printf("Default values: %d, %.2f, %t, '%s'\n",
		defaultInt, defaultFloat, defaultBoolean, defaultString)
	return

	//Which declaration form to use?
	//No short declaration can be used outside of functions including main.

	//Making a variable accessible to other packages, i.e. exporting
	// 1. Only long or multiple declaration
	// 2. Capital first letter

	//Task: Define variables of various types at package level,
	//print them in main function,
	//create a variable with the same naming, different value.
}
