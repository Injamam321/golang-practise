package main

import "fmt"

func main() {
	num := -10
	if num > 0 {
		fmt.Println(num, "is positive")
	} else {
		fmt.Println(num, "is not positive")

	}

	num := 0
	if num > 0 {
		fmt.Println(num, "number is positive")

	} else if num == 0 {
		fmt.Println(num, "number is Zero")
	} else {
		fmt.Println(num, "number is negative")
	}

	// nested if else
	num := 50
	if num > 0 {
		fmt.Println(num, "number is positive")

	} else {
		if num == 0 {
			fmt.Println(num, "number is Zero")
		} else {
			fmt.Println(num, "number is negative")
		}
	}

	//if with function
	num := -12
	if isPositive(num) {
		fmt.Println(num, "number is positive")
	} else {
		fmt.Println(num, "number is Neagtive")
	}

}

func isPositive(num int) bool {
	return num > 0
}
