package main

import "fmt"

// In go, we can declare variables in two ways: using 'var' keyword or using short short variable declaratio.

func main() {
	// First, Let's try 'var' method.
	// We first write variable name, and then its type. We can pmit either the type or the expression	 // in assigning, but not both.
	var first int = 32
	var second int
	var third = 64
	var fifth, sixth = "Hi", 55
	fmt.Printf("first: value = %d, type = %T\n", first, first) // We have printf in Go too :( Thanks God.
	fmt.Printf("third: value = %d, type = %T\n", third, third)
	fmt.Printf("fifth = %s, sixth = %d\n", fifth, sixth)
	fmt.Printf("second = %d\n", second)
	fmt.Println("__________________________________________")

	// The second way is using short variable declaration.
	// note: We can use this way just for local variables.
	var1 := 21
	var2 := "Hello world!"
	var3 := 33.44
	fmt.Printf("var1 = %d, var2 = %s, var3 = %f\n", var1, var2, var3)
	var4, var5, var6 := 4, 5, 6
	fmt.Println(var4, var5, var6)
}
