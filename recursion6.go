package main

import "fmt"

func fact(n int) int {
    if n == 1 {
        return 0
    }
    return n * fact(n-1)
}

func main() {
    fmt.Println(fact(8))
}