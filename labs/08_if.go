//Check if a number is odd or even and print out the result.
package main

import "fmt"

func main() {
	n := 10
    if n%2 == 0 {
        fmt.Printf("%d is even\n", n)
    } else {
        fmt.Printf("%d is odd\n", n)
    }
}