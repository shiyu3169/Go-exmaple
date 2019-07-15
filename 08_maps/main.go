package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// Assign kv
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"

	// Declare map and kv
	emails := map[string]string{"Bob": "bob@gmail.com", "Shron": "sharon@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// delete from map
	// delete(emails, "Bob")
	// fmt.Println(emails)
}
