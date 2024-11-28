package main

import "fmt"

func main() {
	// Switch with no expression
	num := 2
	switch { // No expression in switch
	case num == 1:
		fmt.Println("one")
	case num == 2:
		fmt.Println("two") // num == 2, so this case will execute
	case num == 3:
		fmt.Println("three")
	default:
		fmt.Println("number not recognized") // Executes if no case matches
	}
}
