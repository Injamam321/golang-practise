package main

import "fmt"

func main() {
	// Another fallthrough example

	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Start of the week")
		fallthrough // Continues to the next case, regardless of condition
	case "Tuesday":
		fmt.Println("Second day")
		fallthrough // Continues to the next case, regardless of condition
	case "Wednesday":
		fmt.Println("Midweek")
		fallthrough // Continues to the next case, regardless of condition
	case "Thursday":
		fmt.Println("Almost weekend")
		// No fallthrough needed here, default won't be reached
	default:
		fmt.Println("End of the week")
	}
}
