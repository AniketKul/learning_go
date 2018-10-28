package main

import "fmt"

func main() {
	messagePrinterStruct := messagePrinter{"foo"}
	messagePrinterStruct.printMessage()
}

type messagePrinter struct {
	message string
}

func (mp *messagePrinter) printMessage() {
	fmt.Println(mp.message)
}
