package main

import "fmt"

func main() {
	// Switch with no expression (using multiple conditions)

	age := 20
	switch {
	case age < 18:
		fmt.Println("You are a minor.")

	case age >= 18 && age <= 30:
		fmt.Println("You are a young adult.") // This will execute, because age is 20

	case age > 30:
		fmt.Println("You are an adult.")

	default:
		fmt.Println("Age is not recognized.")
	}
}
