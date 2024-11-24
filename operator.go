package main

import "fmt"

func main() {

	// arithmatic operator
	// a := 10
	// b := 5

	// fmt.Println(a + b)
	// fmt.Println(a - b)
	// fmt.Println(a * b)
	// fmt.Println(a / b)
	// fmt.Println(a % b)

	// increment operator
	// a++
	// b++
	// fmt.Println(a)
	// fmt.Println(b)

	// Decremet operator
	// 	a--
	// 	b--
	// 	fmt.Println(a)
	// 	fmt.Println(b)
	// 	a--
	// 	fmt.Println(a)
	// 	b--
	// 	fmt.Println(b)

	// relational operator

	// ==
	// a := 9
	// b := 6
	// c := 9

	// fmt.Println(a == c)
	// fmt.Println(a == b)

	// !=
	// a := 9
	// b := 6
	// c := 9

	// fmt.Println(a != c)
	// fmt.Println(a != b)

	// < >

	// a := 9
	// b := 6
	// c := 9

	// fmt.Println(a < c)
	// fmt.Println(a < b)
	// fmt.Println(a > b)

	// >=
	// a := 9
	// b := 6
	// c := 9

	// // fmt.Println(a >= c)
	// // fmt.Println(a >= b)
	// // fmt.Println(b >= c)
	// // fmt.Println(b >= a)

	// //<=
	// fmt.Println(a <= b)
	// fmt.Println(a <= c)

	//logical operator

	// logical and &&
	// t := true  // 5==5
	// f := false // 5 > 7
	// f1 := false
	// t1 := true

	// fmt.Println(t1 && t)
	// fmt.Println(f && f1)
	// fmt.Println(t && f)
	// fmt.Println(t1 && f1)

	// logical or ||

	// fmt.Println(t1 || t)
	// fmt.Println(f || f1)
	// fmt.Println(t || f)
	// fmt.Println(t1 || f1)

	// logical not !
	// fmt.Println(!t)
	// fmt.Println(!f1)
	// fmt.Println(!f)
	// fmt.Println(!t1)

	// n-> 5 and 6 visible or not
	// n := 30
	// div5 := n%5 == 0
	// div6 := n%6 == 0
	// result := div5 && div6
	// fmt.Println(result)

	// bitwise operator
	// bitwise AND &
	// bitwise OR
	//bitwise exclusive or ( ^ ) two same value return 0 otherwise 1
	//bitwise not ^ single vaule er upor impact kore

	// a := 5             // 0101
	// b := 3             // 0011
	// fmt.Println(a & b) // & 0001=1
	// fmt.Println(a | b) // 0111=7
	// fmt.Println(a ^ b) // 0110 =6
	// fmt.Println( ^a)    //1010=10

	//assignment operator
	a := 6
	a = 9
	fmt.Println(a) //9

	a += 5
	fmt.Println(a) //14

	a -= 6
	fmt.Println(a) //8

	a *= 5
	fmt.Println(a) //40

	a /= 5
	fmt.Println(a) // 8

	a %= 5
	fmt.Println(a) //3

}
