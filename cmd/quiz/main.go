package main

import (
	"fmt"
	"time"
)

func timer(seconds int) <-chan bool {
	done := make(chan bool)
	// Initiate timer
	go func() {
		time.Sleep(time.Duration(seconds) * time.Second)
		done <- true
	}()
	// Listen for time stoper
	return done
}

// Learn about: flags, CSV, OS, time package.
func main() {
	t := timer(10)
	for {
		status := <-t
		if status == true {
			fmt.Println("done")
			return
		}
	}
}
