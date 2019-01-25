//Slices

package main

import "fmt"

func main() {
	//A slice does not store any data,
	//it just describes a section of an underlying array.
	//The value of an uninitialized slice is nil.
	var s1 []int
	fmt.Println("Empty slice:", s1, "Length:", len(s1))

	//return

	//A slice is a dynamically-sized array.
	//Changing the elements of a slice modifies
	//the corresponding elements of its underlying array.
	s1 = append(s1, 2)
	fmt.Println("Not empty slice", s1[0], len(s1))

	//return

	//Capacity
	fmt.Println("Capacity", cap(s1))

	//What is capacity?
	s1 = append(s1, 3, 4)
	fmt.Println(s1, "Length", len(s1), "Capacity", cap(s1))
	s1 = append(s1, 5, 6)
	fmt.Println(s1, "Length", len(s1), "Capacity", cap(s1))
	s1 = append(s1, 7, 8)
	fmt.Println(s1, "Length", len(s1), "Capacity", cap(s1))
	s1 = append(s1, 9, 10)
	fmt.Println(s1, "Length", len(s1), "Capacity", cap(s1))
	//The number of elements in the underlying array,
	//counting from the first element in the slice.

	//Short declaration
	s2 := []int{10, 20, 30}
	fmt.Println(s2, len(s2), cap(s2))

	//How to add one slice to another.
	//s1 = append(s1, s2)
	//fmt.Println(s1)

	//Slices are one-dimensional. To create the equivalent of a 2D slice,
	//it is necessary to define an slice-of-slices.
	var multi_s [][]int
	multi_s = append(multi_s, s1, s2)
	fmt.Println(multi_s)

	//Declare a slice of defined length
	s3 := make([]int, 5)
	fmt.Println(s3, "Length", len(s3), "Capacity", cap(s3))

	//What will be printed here?
	s3 = append(s3, 1)
	fmt.Println(s3)

	//Declare a slice of given length and capacity
	s4 := make([]int, 5, 10)
	fmt.Println(s4, "Length", len(s4), "Capacity", cap(s4))

	//How to declare a slice, int type, 0 length, capacity 20
	s5 := make([]int, 0, 20)
	fmt.Println(s5, "Length", len(s5), "Capacity", cap(s5))

	//Task: Write a function that calculates the average of a float62 slice.
}
