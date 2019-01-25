//
package main

import "fmt"

func main() {
    for i := -5; i <= 5; i++ {
        if i%2 == 0 {
            fmt.Printf("%d is even\n", i)
        } else {
            fmt.Printf("%d is odd\n", i)
        }
    }
}
