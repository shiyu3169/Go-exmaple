package main

import "fmt"

func main() {
	// Using var

	var age = 27
	// Const
	const isCool = true
	// Shorthand
	name := "Shiyu"

	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n", isCool)
}
