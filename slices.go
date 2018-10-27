package main

import "fmt"

func main() {
	mySlice := make([]float32, 100)
	mySlice[0] = 14.
	mySlice[1] = 15.
	mySlice[2] = 16.
	// mySlice := []float32{14., 15., 16.}
	fmt.Println(mySlice)
	fmt.Println("Length of the slice: ", len(mySlice))
}
