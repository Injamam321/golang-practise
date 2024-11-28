package main

import "fmt"

func main() {

	// multiple expressions in case

	num := 2
	switch num {
	case 1, 2:
		fmt.Println("one or two") // যদি num 1 অথবা 2 হয়
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("number not recognized")
	}

	day := "Friday"
	switch day {
	case "Monday", "Wednesday", "Friday":
		fmt.Println("Start of the week or Mid-week or Near weekend")

	case "Tuesday", "Thursday":
		fmt.Println("Mid-week")

	case "Saturday", "Sunday":
		fmt.Println("Weekend")

	default:
		fmt.Println("Invalid day")
	}
}
