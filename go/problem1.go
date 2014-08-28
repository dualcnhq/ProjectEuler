// Multiples of 3 and 5
// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. 
// The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.

package main

import "fmt"

func is_multiple(i int) bool {
    return i % 5 == 0 || i % 3 == 0;
}

func main() {
    s := 0
    for i := 0; i < 1000; i++ {
        if is_multiple(i) {
            s += i
        }
    }
    fmt.Printf("Sum: %v\n", s)
}
