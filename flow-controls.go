package main

import (
	"fmt"
	"runtime"
)

func main() {
	// for loop in very similar to C:
	end := 4
	for i := 0; i < end; i++ {
		fmt.Println("Viva Go!")
	}

	// init and post statment are optional. So technically we write while loop using for in Go :)) :
	counter := 0
	for counter < 3 {
		fmt.Println("That's how we write while in Go(Sunglasses emoji :))")
		counter++
	}

	// If we don't write loop condition too, we have infinite loop!
	counter = 0
	for {
		if counter < 3 {
			fmt.Println("Not yet!")
			counter++
		} else {
			break
		}
	}
	// We write if like for too. We say they're alike, because we can write a short statement in if too.

	// We have switch in Go too, and it's better than its version in C. Values don't have to be just in or
	// constant numbers, and we have short statement in switch case too:
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("X")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("another boolshit OS like %s! :)", os)

	}
	// We can write switch with no condition, and it's sometimes a really cleaner way to write if-else code:
	name := "Parham"
	switch {
	case name < "Nelin":
		fmt.Println(":))")
	default:
		fmt.Println("logical answer!")
	}

	// We have a new thing in Go called defer. It defers the execution of a function until the surrounding
	// function returns!
	defer fmt.Println("As you can see, this sentence is printed after main function is returned")
	// A very interesting fact is that multiple defers work in stack order, or LIFO:
	for i := 0; i < 3; i++ {
		defer fmt.Printf("print number %d\n", i)
	}
	return
}
