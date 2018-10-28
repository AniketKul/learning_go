package main

import "fmt"

func main() {
	foo := initializeMyStruct()
	foo.myMap["India"] = "Delhi"
	foo.myMap["Korea"] = "Seoul"
	foo.myMap["Japan"] = "Tokyo"

	fmt.Println(foo)
}

type myStruct struct {
	myMap map[string]string
}

func initializeMyStruct() *myStruct {
	result := myStruct{}
	result.myMap = map[string]string{}
	return &result
}
