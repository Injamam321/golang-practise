package main

import "fmt"

func main() {
	// var imu int
	// imu = 5
	// fmt.Println(imu)
	// imu = 7
	// fmt.Println(imu)

	// const imu int
	// imu = 5
	// fmt.Println(imu)
	// imu = 71
	// fmt.Println(imu)
	//iota
	const (
		First = iota
		second
		Third
	)
	fmt.Println(First, second, Third)

}
