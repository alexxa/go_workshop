//Constants

/*
Constants are declared like variables, but with the const keyword.
Constants may be typed or untyped (explicit or implicit)
There are string, boolean, and numeric constants.
Constants cannot be declared using the := syntax.
*/

package main

func main() {

	const pi float64 = 3.14
	println(pi)

	//return

	//Multiple constant declarations
	const (
		x float64 = 1.5
		x2
		y = 0
	)
	println(x, x2, y)

	//return

	//Multiple constants can be defined on the same line
	//as well just as in multiple variables declaration

	//Constants and type convertion
	const a int = 42
	const b = -42
	println(a + b)

	//return

	//Error
	const y1 uint = 10
	//println(x1 + y1)

	const (
		i = iota
		j = iota
	)

	println(i, j)

	const (
		p = iota
		q
	)

	println(p, q)

	//return
}
