package main

func main() {
	message := "Hello Aniket!"
	sayHello1(message)  // Passing by value
	sayHello2(&message) // Passing by reference
	println(message)
	sayHello3("Hello U.S.A", "Hello India", "Hello Japan", "Hello Korea")
}

// Passing by value
func sayHello1(message string) {
	println(message)
	message = "Hello Go!" // did not change the value of message var
}

// Passing by reference
func sayHello2(message *string) {
	println(*message)
	*message = "Hello Go!"
}

// Variadic function a.k.a function that accepts any number of parameters
func sayHello3(messages ...string) {
	// messages is slice here
	for _, message := range messages {
		println(message)
	}
}
