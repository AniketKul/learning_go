package main

//By default go programs only run on single CPU. (Concurrency)
import "time"

// Allows go routines to use multiple CPUs (Parallelism)
import "runtime"

// channels allow applications to use multiple CPUs without messing up the order of tasks.

func main() {
	runtime.GOMAXPROCS(8) // parallel programming
	// Run it co-currently. A bit expensive in terms of memory
	go abcGen()
	println("This comes first!")
	time.Sleep(100 * time.Millisecond)
}

func abcGen() {
	for l := byte('a'); l <= byte('z'); l++ {
		go println(string(l))
	}
}
