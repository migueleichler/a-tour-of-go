package main

import "fmt"


// 1 - A function can take zero or more arguments.
// 2 - The type of the variable comes after the variable name.
func add(x int, y int) int {
    return x + y
}

// 3 - When two or more consecutive named function parameters share a type,
//     you can omit the type from all but the last.
func subtract(x, y int) int {
    return x - y
}

// 4 - A function can return any number of results
func multipleReturnArguments(x, y int) (int, int) {
    return subtract(x, y), subtract(y, x)
}

// 5 - The return values may be named.
// 6 - A return statement without arguments returns the named return values.
//     This is known as a "naked" return.
func withoutReturnArguments(a, b int) (x, y int) {
    x = subtract(a, b)
    y = subtract(b, a)

    return
}

func main() {
    var x int = 2
    var y int = 3

    fmt.Println(add(x, y))
    fmt.Println(subtract(x, y))
    fmt.Println(multipleReturnArguments(x, y))
    fmt.Println(withoutReturnArguments(x, y))
}
