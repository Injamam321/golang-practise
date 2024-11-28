package main

import "fmt"

//global
//local

var globalvar = "global value"

func main() {

	var localvar = "local value" //scope

	fmt.Println(globalvar)
	fmt.Println(localvar)

}
func love() {
	fmt.Println(globalvar)

	//fmt.Println(localvar)
}
