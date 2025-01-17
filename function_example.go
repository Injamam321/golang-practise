//Normal function declare

package main

import "fmt"

func add(a int, b int) int {
	sum := a + b
	return sum
}

func main() {
	result := add(5, 7)
	fmt.Println("Sum:", result)
}


//multiple value return for a function

package main

import "fmt"

func divide(x int, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}

func main() {
	quotient, remainder := divide(10, 3)
	fmt.Println("Quotient:", quotient, "Remainder:", remainder)
}
