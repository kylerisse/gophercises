package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

type question struct {
	question string
	answer   string
}

func main() {
	filePath, timeLimit, isShuffle := cliFlags()
	questions := csvToQuestions(filePath)
	if isShuffle {
		questions = shuffleQuestions(questions)
	}
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
	fmt.Printf("Answer the following questions in %d seconds:\n", timeLimit)
	fmt.Printf("---------------------------------------------\n")
	for i := range questions {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Question #%d: %v = ", i+1, questions[i].question)
		a, _ := reader.ReadString('\n')
		a = strings.Replace(a, "\n", "", -1)
		a = strings.Trim(a, " ")
		a = strings.ToLower(a)
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
		if err == io.EOF {
			break
		}
		if len(record) != 2 {
			continue
		}
		qs = append(qs, question{
			question: record[0],
			answer:   strings.ToLower(record[1]),
		})
	}
	return qs
}

func shuffleQuestions(qs []question) []question {
	for i := range qs {
		j := rand.Intn(i + 1)
		qs[i], qs[j] = qs[j], qs[i]
	}
	return qs
}

func cliFlags() (string, int, bool) {
	var filePath string
	flag.StringVar(&filePath, "f", "problems.csv", "specify problem file.")
	var timeLimit int
	flag.IntVar(&timeLimit, "t", 30, "specify time limit.")
	var shuffle bool
	flag.BoolVar(&shuffle, "s", false, "shuffle questions.")
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
	return filePath, timeLimit, shuffle
}
