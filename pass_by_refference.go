package main

import "fmt"

// Function to demonstrate pass by reference
func passByReference(y *int) {
	*y = 20
}

func main() {
	// Pass by reference
	value := 5
	passByReference(&value) // Pass the address of 'value'
	fmt.Println(value)      // This will print 20 because 'value' is modified through its reference
}
