package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	pointers()
	arrays()
	work_with_slices1()
	work_with_slices2()
	range_form()
	maps()
}

func pointers() {
	// We have pointers in Go too. They're much like C pointers, except that we don't have arithmetic for them
	i := 42
	i_pointer := &i
	fmt.Println(*i_pointer)
	*i_pointer = 30

	fmt.Println(i)
	j := 67.7
	var j_pointer *float64 = &j
	fmt.Println(*j_pointer)

	// Like C, we have structs in Go too!
	v1 := Vertex{11, 15}
	fmt.Println(v1)
	v1.X = 66
	fmt.Println(v1)

	// We have pointer to structs here too! But unlike C, we don't have '->' notation here. We can use dot
	// with pointers without any problems here:
	v_pointer := &v1
	fmt.Println(*v_pointer)
	v_pointer.Y = 88
	fmt.Println(*v_pointer)
}

func arrays() {
	// We have arrays here too we declare them like these:
	var a [3]int
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	fmt.Println(a)
	// or
	primes := [4]int{2, 3, 5, 7}
	fmt.Println(primes)

	// We have slices in go too. We declare that like arrays, but we don't specify size for them in square brackets
	slice1 := primes[1:3]
	fmt.Println(slice1)
	var slice2 []int = a[1:2]
	fmt.Println(slice2)
	// Notice: like python, slices just select a subarray of the original array. They don't store any values.

	// Just a little complex method to create a slice from a struct
	st := []struct {
		X int
		Y int
	}{
		{1, 2},
		{3, 4},
		{9, 10},
		{11, 14},
	}
	fmt.Println(st)
}

func work_with_slices1() {
	slice := []int{1, 2, 5, 7, 9, 111, 2232}
	slice_detail(slice)

	slice = slice[1:3]
	slice_detail(slice)

	slice = slice[:1]
	slice_detail(slice)

	slice = slice[:]
	slice_detail(slice)
}

func work_with_slices2() {
	// In this function, we make slices with built-in make function.
	slice := make([]int, 5)
	slice_detail((slice))

	slice2 := make([]int, 0, 5) // third argument is capacity
	slice_detail(slice2)

	slice3 := slice[:2]
	slice_detail((slice3)) 

	slice4 := slice[2:]
	slice_detail(slice4)

	// We can expand slice size using append function. If the size of the array is not enough,
	// It will allocate a new array and point to that.
	fmt.Println(slice4)
	slice4 = append(slice4, 9, 8, 7)
	slice_detail(slice4)
	fmt.Println(slice4)
}

func range_form() {
	// Range form of a loop iterates over a slice, and returns index and copy of value at that index.
	slice := []int{1, 2, 3, 5, 9, 10}
	for index, value := range slice {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}

	// We can skip index or value using _
	for _, value := range slice {
		fmt.Println(value)
	}

	for index, _ := range slice {
		fmt.Printf("index: %d\n", index)
	}

	// We can omit value if we just want to use index.
	for index := range slice {
		fmt.Println(index)
	}
}

func slice_detail(slice []int) {
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
}

// Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers.
// When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.
func exercise(dx, dy int) [][]uint8 {
	// This function is an execise in A tour of go
	image := make([][]uint8, dx)
	for index1 := range image {
		image[index1] = make([]uint8, dy)
		for index2 := range image[index1] {
			image[index1][index2] = uint8(index1) ^ uint8(index2)
		}
	}	
	return image
}

func maps() {
	// we define maps with this syntax: map[key-type]value-type
	// we use make to make an empty and ready to use map
	a := make(map[string]int)
	a["Parham"] = 19
	a["Nelin"] = 20
	fmt.Println(a["Parham"], a["Nelin"])
	
	// We delete elements with what? YES; delete :)
	delete(a, "Parham")

	// We check existance of element with two-value assignement
	value, is_in := a["Parham"]
	fmt.Println(value, is_in)
	fmt.Println(a)
}