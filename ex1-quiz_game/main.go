package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type question struct {
	question string
	answer   string
}

func main() {
	filePath, timeLimit := cliFlags()
	questions := csvToQuestions(filePath)
	var score int
	go gameLoop(questions, &score, timeLimit)
	time.Sleep(time.Duration(timeLimit) * time.Second)
	gameOver(score, len(questions))
}

func gameOver(score int, total int) {
	fmt.Printf("\nYou scored %d out of %d\n", score, total)
	os.Exit(0)
}

func gameLoop(questions []question, score *int, timeLimit int) {

	fmt.Printf("Answer the following questions n %d seconds:\n", timeLimit)
	fmt.Printf("--------------------------------------------\n")
	for i := range questions {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Question #%d: %v = ", i+1, questions[i].question)
		a, _ := reader.ReadString('\n')
		a = strings.Replace(a, "\n", "", -1)
		if strings.Compare(questions[i].answer, a) == 0 {
			*score++
		}
	}
	gameOver(*score, len(questions))
}

func csvToQuestions(filePath string) []question {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("can't open file %v\n", filePath)
		os.Exit(2)
	}
	r := csv.NewReader(bufio.NewReader(f))
	var qs []question
	for {
		record, err := r.Read()
		if err == io.EOF || len(record) != 2 {
			break
		}
		qs = append(qs, question{
			question: record[0],
			answer:   record[1],
		})
	}
	return qs
}

func cliFlags() (string, int) {
	var filePath string
	flag.StringVar(&filePath, "f", "problems.csv", "specify problem file.")
	var timeLimit int
	flag.IntVar(&timeLimit, "t", 30, "specify time limit.")
	var help bool
	flag.BoolVar(&help, "h", false, "display this message.")
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if help == true {
		flag.Usage()
		os.Exit(0)
	}
	if timeLimit < 1 {
		fmt.Printf("Invalid time limit %d\n", timeLimit)
		os.Exit(1)
	}
	return filePath, timeLimit
}
