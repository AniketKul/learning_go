package main

func main() {
	for i := 0; i < 5; i++ {
		println(i)
	}

	println("2nd time of for loop: ")
	i := 1

	for {
		i++
		println(i)

		if i > 5 {
			break
		}
	}

	// traversing through slice
	s := []string{"foo", "bar", "buz"}

	for index := 0; index < len(s); index++ {
		println(index, s[index])
	}

	for idx, v := range s {
		println(idx, v)
	}

	// traversing through map
	myMap := make(map[string]string)
	myMap["first"] = "first"
	myMap["second"] = "second"
	myMap["third"] = "third"

	for k, v := range myMap {
		println(k, v)
	}

}
