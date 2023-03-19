package main

import "fmt"

func main() {
	// We create maps in three different ways

	// #1
	contacts1 := map[string]string{
		"Parham": "P_roufarshbaf@yahoo.com",
		"Amirali": "A_roufarshbaf@yahoo.com",
		"Gholi"  : "Gholi@yahoo.com",
	}

	// #2
	var contacts2 map[string]string

	// #3
	contacts3 := make(map[string]string)

	// deleting
	delete(contacts1, "Parham")
	fmt.Println(contacts2)
	fmt.Println(contacts1)
	fmt.Println(contacts3)
	printMap(contacts1)
}

func printMap(c map[string]string){
	for name, email := range c {
		fmt.Println("The user with name", name, "has email with address", email)
	}
}