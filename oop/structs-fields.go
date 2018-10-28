package main

import "fmt"

func main() {
	foo := myStruct{"bar"}
	// foo := myStruct{}
	// foo := new(myStruct)
	// foo := &myStruct{"bar"}
	// foo.myField1 = "bar"
	fmt.Println(foo)
}

type myStruct struct {
	myField1 string
}
