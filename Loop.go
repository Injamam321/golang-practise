// foor loop example

// package main

// import "fmt"

// func main() {
// 	for i := 1; i <= 5; i++ {
// 		fmt.Println("Number:", i)
// 	}
// }

// Like while loop but its for loop
// package main

// import "fmt"

// func main() {
// 	num := 1
// 	for num <= 5 {
// 		fmt.Println("Counting:", num)
// 		num++
// 	}
// }

// Infinite loop

package main

import "fmt"

func main() {
	for {
		fmt.Println("This is an infinite loop.")
		break // Break stops the loop
	}
}
