package main

import "fmt"

func main() {
	add := func(a, b int) int {
		return a + b
	}

	fmt.Println("10 + 20 =", add(10, 20))
}
