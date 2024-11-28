package main

import "fmt"

func main() {
	result := func(a, b int) int {
		return a + b
	}(10, 20)

	fmt.Println("Sum:", result)
}
