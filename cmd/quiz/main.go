package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path"
	// "time"
)

// Error interface
type Error interface {
	Error() string
}

// func timer(seconds int) <-chan bool {
//  done := make(chan bool)
//  // Initiate timer
//  go func() {
//      time.Sleep(time.Duration(seconds) * time.Second)
//      done <- true
//  }()
//  return done
// }

func getFile(s string) (*os.File, Error) {
	dir := path.Join("/Users/ericyang/Documents/Home/go/src/github.com/ericyang321/godroplet/cmd/quiz", s)
	return os.Open(dir)
}

// Learn about: flags, CSV, OS, time package.
func main() {
	f, openErr := getFile("problems.csv")
	if openErr != nil {
		fmt.Println(openErr.Error())
		return
	}
	defer f.Close()
	reader := csv.NewReader(f)
	var (
		record   []string
		readErr  Error
		question string
		answer   string
		score    int
	)
	for {
		record, readErr = reader.Read()
		if readErr == io.EOF {
			fmt.Println("End of Quiz! You scored", score)
			return
		} else if readErr != nil {
			fmt.Println("Improper CSV Format", readErr)
			return
		}
		question = record[0]
		answer = record[1]

		fmt.Println(question, answer)
	}
}
