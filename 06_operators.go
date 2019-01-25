/*
Operators: arithmetic, relational, logical, assignment.

Arithmetic:
Operator 	Name 			Types
+ 			sum 	    	integers, floats, complex values, strings
- 			difference 		integers, floats, complex values
* 			product
/ 			quotient
% 			remainder 		integers
& 			bitwise AND
| 			bitwise OR
^ 			bitwise XOR
&^ 			bit clear (AND NOT)
<< 			left shift 		integer << unsigned integer
>> 			right shift 	integer >> unsigned integer

Relational (comparison)
Operator 	Name 			Types
== 			equal 			comparable
!= 			not equal
< 			less 			integers, floats, strings
<= 			less or equal
> 			greater
>= 			greater or equal

Logical
Operator 	Name 				Description
&& 			conditional AND 	a && b   means "if a then b else false"
|| 			conditional OR 		a || b   means "if a then true else b"
! 			NOT 				!a   means "not a"

*/

package main

import "fmt"

func main() {

	var x1 int = 100
	var x2 int = 25
	var x3 int = 5
	var x4 int

	fmt.Println(x1, x2, x3, x4)

	/*
		Complete code which prints results of
		x1 / x2 - x3
		x1 + x2 * x3
		x1 % x2
		x1++
		x1 % x2
	*/

	/*
		Arithmetic operators and result type
		- If both operands have the same type, the result is of the same type
		- If one operand is untyped, and the second is typed, the result is ...
		- If two operands have different type, the result value is the latest type
	*/

	//Write an example demonstrating it

	const X, Y, Z = 1, '*', 2i // three untyped values.
	// Default types: int, rune, complex64.

	var a, b int = X, Y // two typed values.

	t1 := X + Y // the type of t1 is the default type of Y: rune (int32).
	t2 := Y - a // the type of t2 is the type of a: int.
	t3 := a * b // the type of t3 is the types of a and b: int.
	t4 := Z * Y // the type of t4 is the default type of Z: complex64.

	fmt.Println(X, Y, Z)        // 1 65 (0+2i)
	fmt.Println(t1, t2, t3, t4) // 66 64 65 (0+130i)
}
