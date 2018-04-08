package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

func terminate(s string) {
	fmt.Println(s)
	os.Exit(1)
}

func getFile(fileName string) *os.File {
	// Any flag.String lists will throw
	file, err := os.Open(fileName)
	if err != nil {
		terminate(fmt.Sprintf("Failed to open CSV file: %s\n", fileName))
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

func funnelUserAnswer(answerCh chan string) {
	var userAnswer string
	// Answer needs to be a pointer address because fmt.Scanf does
	// direct assignment on user input. We're not in dynamic type land anymore...
	fmt.Scanf("%s \n", &userAnswer)
	answerCh <- userAnswer
}

// Learn about: flags, CSV, OS, time package.
func main() {
	fileName := flag.String("csv", "problems.csv", "a CSV in the format of question, answer")
	timeLimit := flag.Int("limit", 30, "Quiz time limit in seconds")
	flag.Parse()

	csvFile := getFile(*fileName)

	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	lines, err := reader.ReadAll()
	if err != nil {
		terminate(fmt.Sprintf("Failed to parse the CSV file: %s \n", err.Error()))
	}
	parsedProblems := parseLines(lines)

	score := 0

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, entry := range parsedProblems {
		fmt.Printf("Problem #%d: %s = ? \n", i+1, entry.question)
		answerCh := make(chan string)
		go funnelUserAnswer(answerCh)

		select {
		case <-timer.C:
			terminate(fmt.Sprintf("\nTime's up! You scored %d out of %d \n", score, len(parsedProblems)))
		case answer := <-answerCh:
			if answer == entry.answer {
				score++
			}
		}
	}
	fmt.Printf("You scored %d out of %d \n", score, len(parsedProblems))
}
