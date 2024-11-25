package main

import "fmt"

func main() {

//basic switch

num := 2
switch num {
case 1:
	fmt.Println("one")
case 2:
	fmt.Println("two")
case 3:
	fmt.Println("three")
default:
	fmt.Println("number not recognized")
}

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

// switch with no expression

// switch without expression

num := 2
switch { // এখানে কোন এক্সপ্রেশন নেই
case num == 1:
	fmt.Println("one")
case num == 2:
	fmt.Println("two") // num 2 হওয়ায় এখানে যাবে
case num == 3:
	fmt.Println("three") //num 3 to go here
default:
	fmt.Println("number not recognized") //if others number to go here
}

package main

import "fmt"

func main() {
	// switch with no expression (using multiple conditions)

	age := 20
	switch {
	case age < 18:
		fmt.Println("You are a minor.")

	case age >= 18 && age <= 30:
		fmt.Println("You are a young adult.") // এখানে যাবে, কারণ age 20

	case age > 30:
		fmt.Println("You are an adult.")

	default:
		fmt.Println("Age is not recognized.")
	}

//Golang-এ fallthrough কিওয়ার্ডটি switch কেসের মধ্যে ব্যবহার হয় যখন তুমি চাও একটি কেসের পরবর্তী কেসটি চালু (execute) হোক, যদিও সেই পরবর্তী কেসের শর্ত পূর্ণ না হয়। অর্থাৎ, একটি কেসের কাজ শেষ হওয়ার পরও পরবর্তী কেসের কোড কার্যকর হবে, যদি সেখানে fallthrough ব্যবহার করা হয়।

//এটি সাধারণত তখন ব্যবহার করা হয়, যখন তুমি চাইবে একাধিক কেস একে একে কার্যকর হোক, এমনকি তাদের শর্ত মেলেও না। তবে, মনে রাখতে হবে যে fallthrough কেবল পরবর্তী কেসটি চালু করবে এবং তার শর্তটি পরীক্ষা করবে না।

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
package main

import "fmt"

func main() {
	// Another fallthrough example

	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Start of the week")
		fallthrough
	case "Tuesday":
		fmt.Println("Second day")
		fallthrough
	case "Wednesday":
		fmt.Println("Midweek")
		fallthrough
	case "Thursday":
		fmt.Println("Almost weekend")
	default:
		fmt.Println("End of the week")
	}
}
