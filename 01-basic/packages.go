// Programs start running in package main.
// An import path is a string that uniquely identifies a package.
// By convention, the package name is the same as the last element of the import path.
// A package's import path corresponds to its location inside a workspace or
// in a remote repository.
package main

import (
    "fmt"
    "math"
)

// You can also write multiple import statements, like: below
// import "fmt"
// import "math"

func main() {
// In Go, a name is exported if it begins with a capital letter.
    fmt.Printf("2³ = %.1f\n", math.Pow(2,3))
}
