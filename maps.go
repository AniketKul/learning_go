package main

import (
	"fmt"
)

func main() {
	myMap := make(map[int]string)
	fmt.Println(myMap)
	myMap[42] = "Aniket"
	myMap[12] = "Foo"
	myMap[34] = "Bar"
	fmt.Println(myMap[89]) // empty string
	fmt.Println(myMap)
}
