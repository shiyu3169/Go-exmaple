package main

import "fmt"

func main() {
	// Using var
	var age = 27

	// Const
	const isCool = true

	// Shorthand
	name := "Shiyu"

	// Float
	size := 1.3

	fmt.Println(name, age, isCool, size)

	// Print type
	fmt.Printf("%T\n", isCool)
}
