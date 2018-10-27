package main

func main() {
	// Named Return Parameters
	addFunc := func(terms ...int) (totalTerms int, sum int) {
		for _, term := range terms {
			sum += term
		}
		totalTerms = len(terms)
		return
	}

	totalTerms, sum := addFunc(1, 2, 3, 4, 5)
	println(totalTerms, sum)
}
