// package main

// import "fmt"

// // দুটি সংখ্যা যোগফল এবং গুণফল বের করার জন্য ফাংশন
// func calculate(num1 int, num2 int) (int, int) {
// 	sum := num1 + num2
// 	multiply := num1 * num2
// 	return sum, multiply // দুটি মান রিটার্ন করা হচ্ছে
// }

// func main() {
// 	addition, multiplication := calculate(5, 3) // ফাংশন থেকে দুটি মান রিটার্ন
// 	fmt.Println("Addition:", addition)         // আউটপুট: 8
// 	fmt.Println("Multiplication:", multiplication) // আউটপুট: 15
// }

// package main

// import "fmt"

// // দুটি সংখ্যা যোগফল এবং গুণফল বের করার জন্য ফাংশন
// func calculate(num1 int, num2 int) (int, int) {
// 	sum := num1 + num2
// 	multiply := num1 * num2
// 	return sum, multiply // দুটি মান রিটার্ন করা হচ্ছে
// }

// func main() {
// 	addition, multiplication := calculate(5, 3)    // ফাংশন থেকে দুটি মান রিটার্ন
// 	fmt.Println("Addition:", addition)             // আউটপুট: 8
// 	fmt.Println("Multiplication:", multiplication) // আউটপুট: 15
// }

package main

import "fmt"

// একটি ফাংশন যা একটি বার্তা প্রিন্ট করে, কোনো মান রিটার্ন করে না
func greet(name string) {
	fmt.Printf("Hello, %s! Welcome to Go.\n", name)
}

func main() {
	greet("Rahim") // ফাংশন কল
}

// package main

// import "fmt"

// // ১. প্যারামিটার এবং রিটার্ন ভ্যালু সহ ফাংশন
// func add(a int, b int) int {
// 	return a + b
// }

// // ২. মাল্টিপল রিটার্ন ভ্যালু সহ ফাংশন
// func calculate(a int, b int) (int, int) {
// 	return a + b, a * b
// }

// // ৩. কোনো রিটার্ন ছাড়া ফাংশন
// func greet(name string) {
// 	fmt.Printf("Hello, %s!\n", name)
// }

// func main() {
// 	// ফাংশন কল
// 	fmt.Println("Addition:", add(5, 10))          // আউটপুট: Addition: 15
// 	sum, product := calculate(3, 4)               // দুটি মান রিটার্ন
// 	fmt.Println("Sum:", sum, "Product:", product) // আউটপুট: Sum: 7 Product: 12
// 	greet("Rahim")                                // বার্তা প্রিন্ট
// }
