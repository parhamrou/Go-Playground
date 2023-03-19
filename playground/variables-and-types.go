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

	// In Go, inlike see we have to do type conversions explicitly. We do this like in python:
	First := 56.9
	Second := uint(First)
	Third := float64(Second) + 0.9
	fmt.Println(First, Second, Third)

	// When on the right side we have a typed value, the ledt side variable takes that type too.
	// But when we have a constant numeric value, the variable takes int, float64 or complex128 type.
	var i = 45
	x := 45.6
	z := 78 + 0.9i
	fmt.Printf("%T, %T, %T\n", i, x, z)

	// We have a 'const' keyword for constant values. The important point is that we can't use ':=' with that. We have to use simple assignement.
	const Hello = "Hello"
	const Pi = 3.14
	fmt.Println(Hello, Pi)
}
