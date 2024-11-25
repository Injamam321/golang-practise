package main

import "fmt"

func main() {
	// Simple for loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}

package main

import "fmt"

func main() {
	// For loop as a while loop
	i := 1
	for i <= 5 { // condition only
		fmt.Println(i)
		i++ // increment i
	}
}

package main

import "fmt"

func main() {
	// Infinite loop
	for {
		fmt.Println("This loop will run forever!")
	}
}

package main

import "fmt"

func main() {
//Loop with break and continue
for i := 1; i <= 10; i++ {
	if i == 5 {
		continue // যখন i 5 হবে, তখন পরবর্তী iteration এ চলে যাবে
	}
	if i == 8 {
		break // যখন i 8 হবে, loop শেষ হয়ে যাবে
	}
	fmt.Println(i)
}

//for loop with range

package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50} // Ekta slice

	fmt.Println("Slice er protiti element er upor cholchi:")
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}

package main

import "fmt"

func main() {
	students := map[string]int{"Rahim": 85, "Karim": 90, "Nabila": 88} // Ekta map

	fmt.Println("Map er protiti key-value pair er upor cholchi:")
	for name, score := range students {
		fmt.Printf("Name: %s, Score: %d\n", name, score)
	}
}

package main

import "fmt"

func main() {
    numbers := []int{5, 10, 15}

    fmt.Println("Shudhu value print korchi:")
    for _, value := range numbers {
        fmt.Println(value)
    }
}

