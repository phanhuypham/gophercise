package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"

	// "strings"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problem.csv", "problem...")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	time := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	fmt.Print(*timeLimit)
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprint("Cannot open the file: %s \n", *csvFileName))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit("Cannot read the file")
	}

	problems := parseLines(lines)

	correct := 0
problemLoop:
	for i, problem := range problems {
		fmt.Printf("Problem %d: %s = \n", i+1, problem.q)
		answerChan := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChan <- answer
		}()
		select {
		case <-time.C:
			fmt.Println()
			break problemLoop
		case answer := <- answerChan:
			if answer == problem.a {
				correct++
			}
		}

	}
	fmt.Printf("Num correct: %d / %d \n", correct, len(problems))

}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}
