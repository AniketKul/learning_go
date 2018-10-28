package main

func main() {
	emp := enhancedMessagePrinter{}
	// emp := enhancedMessagePrinter{messagePrinter{"foo"}}
	emp.message = "foo"
	emp.printMessage()
}

type messagePrinter struct {
	message string
}

func (mp *messagePrinter) printMessage() {
	println(mp.message)
}

type enhancedMessagePrinter struct {
	messagePrinter
}
