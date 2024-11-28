package main

import "fmt"

func passByValue(x int) {
	x = 10
}

func main() {
	//pass by value
	value := 5
	passByValue(value)
	fmt.Println(value)
}
