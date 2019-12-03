package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	problemsFile, err := os.Open(*csvFilename)
	if err != nil {
		log.Fatalln("Can't open file '", *csvFilename, "' because of error:", err)
	}

	csvReader := csv.NewReader(problemsFile)
	lines, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln("Can't read file:", err)
	}
	problems := parseLines(lines)
	correctAnswers := 0

	// As I understand it is a proper way to read a line from stdin
	scanner := bufio.NewScanner(os.Stdin)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

questionLoop:
	for i, problem := range problems {
		fmt.Printf("Problem #%2d: %s\n", i+1, problem.question)
		answerChannel := make(chan string)

		go func() {
			if scanner.Scan() {
				givenAnswer := scanner.Text()
				answerChannel <- givenAnswer
			} else if scanner.Err() != nil {
				log.Fatalln("Can't read from stdin:", scanner.Err())
			}
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break questionLoop
		case givenAnswer := <-answerChannel:
			if givenAnswer == problem.answer {
				correctAnswers++
				fmt.Println("Right answer.")
			} else {
				fmt.Printf("Wrong! Correct answer: %s and yours answer: %s\n", problem.answer, givenAnswer)
			}
		}
	}

	fmt.Printf("You've got %d right answers out of %d\n", correctAnswers, len(problems))
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

type problem struct {
	question, answer string
}
