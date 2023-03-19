package main

import "fmt"

type factory interface {
	getDetail() string
}

// 
type cheeseFactory struct {
	productName string
	quantity int
}

type drinkFactory struct {
	productName string
	quantity int
}

func (cf cheeseFactory) getDetail() string {
	return cf.productName
}

func (df drinkFactory) getDetail() string {
	return df.productName
}

func printDetail(f factory) {
	fmt.Println(f.getDetail())
}

func main() {
	cf := cheeseFactory{"cheese", 10}
	df := drinkFactory{"drink", 11}
	printDetail(cf)
	printDetail(df)
}