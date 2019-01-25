//Maps.

package main

import "fmt"

func main() {
	//A map maps keys to values.
	//The zero value of a map is nil.
	var monthOrder map[string]int
	fmt.Println(monthOrder)

	//A nil map has no keys, nor can keys be added.
	//monthOrder["June"] = 6

	//Use make to initialize a map.
	monthOrder = make(map[string]int)

	monthOrder["June"] = 6
	fmt.Println(monthOrder)

	//Task. Fill the map with for loop
	//Print how many days in the current month.

	//Q: What does v, ok = monthOrder['Monday'] do?

	//Q: How to delete an element from a map?

}
