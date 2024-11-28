package main

import "fmt"

func main() {
	// fallthrough example

	num := 2
	switch num {
	case 1:

		fmt.Println("one")
		fallthrough // fallthrough পরবর্তী কেস চালু করবে

	case 2:

		fmt.Println("two") // এটি এখানে হবে কারণ fallthrough আছে
		fallthrough

	case 3:

		fmt.Println("three") // fallthrough এর কারণে এটি এখানে হবে

	default:

		fmt.Println("number not recognized")
	}
}
