package main

func main() {

	if foo := 2; foo == 1 {
		println("bar")
	} else {
		println("buz")
	}

	switch foo := "this"; foo {
	case "this":
		println("one")
	case "that":
		println("two")
	case "tat":
		println("three")
	case "pat":
		println("four")
	}

	foo := 1

	switch {
	case foo == 1:
		println("one")
	}

}
