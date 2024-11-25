// package main

// import "fmt"

// // ফাংশন যা একটি সংখ্যা দ্বিগুণ করে
// func double(number int) int {
// 	return number * 2
// }

// func main() {
// 	result := double(5) // ফাংশন কল করা হচ্ছে
// 	fmt.Println(result) // আউটপুট: 10
// }

// package main

// import "fmt"

// // একটি ফাংশন যা শুধু একটি মেসেজ প্রিন্ট করে
// func greet() {
// 	fmt.Println("Good Morning! How are you?")
// }

// func main() {
// 	greet() // ফাংশন কল
// }

// package main

// import "fmt"

// // দুটি সংখ্যার যোগফল বের করার জন্য ফাংশন
// func add(num1 int, num2 int) int {
// 	return num1 + num2
// }

// func main() {
// 	sum := add(10, 20) // ফাংশন কল
// 	fmt.Println(sum)   // আউটপুট: 30
// }

// package main

// import "fmt"

// // একটি ফাংশন যা কোনো মান রিটার্ন করে না
// func sayHello(name string) {
// 	fmt.Printf("Hello, %s!\n", name)
// }

// func main() {
// 	sayHello("Rahim") // ফাংশন কল
// }

package main

import "fmt"

// মেসেজ প্রিন্ট করার জন্য একটি ফাংশন
func greet() {
	fmt.Println("Welcome!")
}

// একটি সংখ্যা দ্বিগুণ করার জন্য ফাংশন
func double(number int) int {
	return number * 2
}

// দুটি সংখ্যার যোগফল বের করার জন্য ফাংশন
func add(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	greet()                  // ফাংশন কল
	fmt.Println(double(5))   // আউটপুট: 10
	fmt.Println(add(10, 20)) // আউটপুট: 30
}
