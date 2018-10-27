package main

const (
	cFirst = 1 << (10 * iota)
	cSecond
	cThrird
)

func main() {
	println(cFirst)
	println(cSecond)
	println(cThrird)
}
