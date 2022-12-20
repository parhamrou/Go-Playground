package main

import "fmt"

// Appearently, in Go we write types after variable names, becuase of some good reasons :) :
func times_two(x int) int {
	return x * 2
}

func adder(x, y int) int { // we can omit the type but the last
	return x + y
}

// In go, functions can return multiple values like python:
func swap_string(x, y string) (string, string) {
	return y, x
}

// In Go, return values can have name. A return value without argumenets is called naked return
// and it returns named values.

func named_returns(x, y int) (u, z int) {
	u = x * 2
	z = y * 8
	return
}

func main() {
	fmt.Println(times_two(22))
	fmt.Println(adder(33, 45))
	fmt.Println(swap_string("Hi", "Bye"))
	fmt.Println(named_returns(1, 4))
}
