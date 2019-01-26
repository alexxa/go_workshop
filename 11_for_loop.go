//Loops. For loop. Nested loop.
//Loop control (break, continue, goto).
//Infinite loop.

package main

import "fmt"

func main() {
	//Go has only one looping construct, three-component for loop
	//init and post statements, condition expression
	sum := 0
	for i := 0; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum)
	//return

	//Single condition or while loop
	num := 1
	for num < 10 { //Semicolons are dropped, but can be present
		num *= 2
	}
	fmt.Println(num)

	//Task_02: write infinite loop counting stars

	//Break, break with a label and continue

J:	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break J
			}
			println(i)
		}
	}
}
