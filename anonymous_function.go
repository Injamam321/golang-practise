package main

import "fmt"

func main() {
	func() {
		fmt.Println("hello")
	}() // Adding () to invoke the anonymous function
}
