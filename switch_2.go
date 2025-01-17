package main

import "fmt"

func main() {
	day := "Saturday"

	switch day {
	case "Saturday", "Sunday":
		fmt.Println("It's a weekend!")
	default:
		fmt.Println("It's a weekday.")
	}
}
