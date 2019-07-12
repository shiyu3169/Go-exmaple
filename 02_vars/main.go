package main

import "fmt"

func main() {
	// Using var
	var name = "Shiyu"
	var age = 27
	const isCool = true

	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n", isCool)
}
