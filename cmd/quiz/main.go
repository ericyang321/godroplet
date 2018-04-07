package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Error interface
type Error interface {
	Error() string
}

type problem struct {
	question string
	answer   string
}

// func timer(seconds int) <-chan bool {
// 	done := make(chan bool)
// 	// Initiate timer
// 	go func() {
// 		time.Sleep(time.Duration(seconds) * time.Second)
// 		done <- true
// 	}()
// 	return done
// }

func terminate(s string) {
	fmt.Println(s)
	os.Exit(1)
}

func getFile() *os.File {
	// dir := path.Join("/Users/ericyang/Documents/Home/go/src/github.com/ericyang321/godroplet/cmd/quiz", s)
	// Any flag.String lists will throw
	fileName := flag.String("csv", "problems.csv", "a CSV in the format of question, answer")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		terminate(fmt.Sprintf("Failed to open CSV file: %s\n", *fileName))
	}
	return file
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return problems
}

func checkAnswer(userAnswer string, realAnswer string) bool {
	if userAnswer == realAnswer {
		return true
	}
	return false
}

// Learn about: flags, CSV, OS, time package.
func main() {
	csvFile := getFile()
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	lines, err := reader.ReadAll()
	if err != nil {
		terminate(fmt.Sprintf("Failed to parse theCSV file: %s \n", err.Error()))
	}
	parsedProblems := parseLines(lines)

	score := 0
	for i, entry := range parsedProblems {
		var userAnswer string
		fmt.Printf("Problem #%d: %s = ? \n", i+1, entry.question)
		// Answer needs to be a pointer address because fmt.Scanf does
		// direct assignment on user input. We're not in dynamic type land anymore...
		fmt.Scanf("%s \n", &userAnswer)
		isCorrect := checkAnswer(userAnswer, entry.answer)
		if isCorrect == true {
			score++
		}
	}

	fmt.Printf("You scored %d out of %d \n", score, len(parsedProblems))
}
