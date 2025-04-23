// // package main

// // import "fmt"

// // func main() {
// // 	num := 5

// // 	if num > 0 {
// // 		fmt.Println("Positive Number")
// // 	} else if num < 0 {
// // 		fmt.Println("Negative Number")
// // 	} else {
// // 		fmt.Println("Zero")
// // 	}
// // }

// package main

// import "fmt"

// func main() {
// 	// age := 17

// 	// if age >= 18 {
// 	// 	fmt.Println("You can vote!")
// 	// } else {
// 	// 	fmt.Println("You are too young to vote.")
// 	// }

// 	day := "Monday"

// 	switch day {
// 	case "Monday":
// 		fmt.Println("Start of the week")
// 	case "Friday":
// 		fmt.Println("Weekend is near")
// 	case "Saturday", "Sunday":
// 		fmt.Println("It's the weekend!")
// 	default:
// 		fmt.Println("Midweek day")
// 	}

// }

package main

import "fmt"

func main() {
	marks := 85

	switch {
	case marks >= 90:
		fmt.Println("Grade: A+")
	case marks >= 80:
		fmt.Println("Grade: A")
	case marks >= 70:
		fmt.Println("Grade: B")
	case marks >= 60:
		fmt.Println("Grade: C")
	default:
		fmt.Println("Grade: F")
	}

}
