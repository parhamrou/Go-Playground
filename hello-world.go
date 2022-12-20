package main

// Comments come with '//', like C :)
// Every Go program is made up of some packages. We import packages like this:
import (
	"fmt"
	"math"
)

// we define functions in Go using 'func' keyword.
func main() {
	fmt.Println("Hello world! :)")
	fmt.Println("Why don't we have dear semicolon here?! :(")
	fmt.Println(math.Pi) // Exported names begin with CAPITAL letters in go
}
