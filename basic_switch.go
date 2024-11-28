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
}

//Golang-এ fallthrough কিওয়ার্ডটি switch কেসের মধ্যে ব্যবহার হয় যখন তুমি চাও একটি কেসের পরবর্তী কেসটি চালু (execute) হোক, যদিও সেই পরবর্তী কেসের শর্ত পূর্ণ না হয়।
// অর্থাৎ, একটি কেসের কাজ শেষ হওয়ার পরও পরবর্তী কেসের কোড কার্যকর হবে, যদি সেখানে fallthrough ব্যবহার করা হয়।

//এটি সাধারণত তখন ব্যবহার করা হয়, যখন তুমি চাইবে একাধিক কেস একে একে কার্যকর হোক, এমনকি তাদের শর্ত মেলেও না।
//  তবে, মনে রাখতে হবে যে fallthrough কেবল পরবর্তী কেসটি চালু করবে এবং তার শর্তটি পরীক্ষা করবে না।
