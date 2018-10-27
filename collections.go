package main

import "fmt"

func main() {

	// Arrays are not commonly used in Go. Instead Slices are used.

	myArr := [3]int{}
	// myArr := [...]int{42, 43, 44}
	myArr[0] = 42
	myArr[1] = 43
	myArr[2] = 44

	fmt.Println(myArr)
	fmt.Println("Length of the array: ", len(myArr))

	// Slices: subset of an array
	mySlice := myArr[:]
	mySlice = append(mySlice, 100)
	fmt.Println(myArr)
	fmt.Println(mySlice)

}
