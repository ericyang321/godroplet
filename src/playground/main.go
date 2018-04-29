package main

import (
	"runtime"
	"time"
)

var numCPU = runtime.NumCPU()

func doThing(list []int) {
	sixSeconds := 6 * time.Second
	time.Sleep(sixSeconds)
}

func main() {
	// activity := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// for i := 0; i < numCPU; i++ {

	// }
}
