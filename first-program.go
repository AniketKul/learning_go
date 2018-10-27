package main

func main() {
	var intVar int
	intVar = 43
	println(intVar)

	var floatVar32 float32 = 42.
	println(floatVar32)

	myString := "Hello Go"
	println(myString)

	myComplex := complex(2, 3)
	println(myComplex)
	println(real(myComplex))
	println(imag(myComplex))
}
