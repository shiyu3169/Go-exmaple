package main

import "fmt"

func main() {
	// Using var
	var age = 27

	// Const
	const isCool = true

	// Shorthand
	// name := "Shiyu"
	// email:="brad@gmail.com"

	// shorter
	name, email := "shiyu", "shiyu@gmail.com"

	// Float
	size := 1.3

	fmt.Println(name, age, isCool, size)

	// Print type
	fmt.Printf("%T\n", isCool)
}
