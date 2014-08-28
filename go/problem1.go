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
