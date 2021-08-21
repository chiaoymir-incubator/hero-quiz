package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Quiz Game Start!")

	fileName := flag.String("csv", "./problems/main.csv", "the input csv file, the format is question,answer.")
	timeLimit := flag.Int("limit", 180, "the question time limit")
	quizNumber := flag.Int("num", 10, "the number of question")
	flag.Parse()

	if *timeLimit < 1 {
		colorizeln(ColorRed, "The time should be positive")
		return
	}

	file := readCSV(fileName)
	quizs := parseCSVQuiz(file)

	startQuiz(quizs, *timeLimit, *quizNumber)

}
