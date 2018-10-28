package main

// By default go programs only run on single CPU. (Concurrency)
// import "time"

// Allows go routines to use multiple CPUs (Parallelism)
import "runtime"

func main() {
	runtime.GOMAXPROCS(8) // parallel programming

	// channels allow applications to use multiple CPUs without messing up the order of tasks.
	ch := make(chan string)
	doneCh := make(chan bool)
	// Run it co-currently. A bit expensive in terms of memory
	go abcGen(ch)
	go printer(ch, doneCh)
	println("This comes first!")
	<-doneCh
}

func abcGen(ch chan string) {
	for l := byte('a'); l <= byte('z'); l++ {
		ch <- string(l) // <- is called receive operator. Used by channels in Go
	}
	close(ch)
}

func printer(ch chan string, doneCh chan bool) {
	for l := range ch {
		println(l)
	}
	doneCh <- true
}
