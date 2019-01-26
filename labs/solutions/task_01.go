//Result type on arithmetic operations.
program task01

import "fmt"

func task01() {

	const x, y, z = 1, '*', 2i // three untyped values.
	// Default types: int, rune, complex64.

	var a, b int = x, y // two typed values.

	t1 := x + y // the type of t1 is the default type of Y: rune (int32).
	fmt.Printf("%T\n", t1)
	t2 := y - a // the type of t2 is the type of a: int.
	fmt.Printf("%T\n", t2)
	t3 := a * b // the type of t3 is the types of a and b: int.
	fmt.Printf("%T\n", t3)
	t4 := z * y // the type of t4 is the default type of Z: complex64.
	fmt.Printf("%T\n", t4)

	fmt.Println(X, Y, Z)
	fmt.Println(t1, t2, t3, t4)
}