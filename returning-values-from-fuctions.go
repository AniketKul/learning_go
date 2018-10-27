package main

// Single Value Return
// Multiple Value Return
// Named Return Parameters

func main() {
	result := addFunc1(1, 2, 3, 4, 5)
	println(result)
	totalTerms, result := addFunc2(1, 2, 3, 4, 5)
	println(totalTerms, result)
	totalTerms, sum := addFunc2(1, 2, 3, 4, 5)
	println(totalTerms, sum)
}

// Single Value Return
func addFunc1(terms ...int) int {
	result := 0
	for _, term := range terms {
		result += term
	}
	return result
}

// Multiple Value Return
func addFunc2(terms ...int) (int, int) {
	result := 0
	for _, term := range terms {
		result += term
	}
	return len(terms), result
}

// Named Return Parameters
func addFunc3(terms ...int) (totalTerms int, sum int) {
	for _, term := range terms {
		sum += term
	}
	totalTerms = len(terms)
	return
}
