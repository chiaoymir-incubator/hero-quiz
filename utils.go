package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func colorize(color Color, message string) {
	fmt.Printf("%s%s%s\n", string(color), message, string(ColorReset))
}

func PrintResult(totalQ, correct, pass, wrong int) {
	fmt.Println("")
	fmt.Println("Result:")

	totalQMsg := fmt.Sprintf("  Number of Questions: %d", totalQ)
	colorize(ColorBrightWhite, totalQMsg)

	correctMsg := fmt.Sprintf("  Correct: %d", correct)
	wrongMsg := fmt.Sprintf("  Wrong: %d", wrong)
	passMsg := fmt.Sprintf("  Pass: %d", pass)
	colorize(ColorBrightGreen, correctMsg)
	colorize(ColorBrightRed, wrongMsg)
	colorize(ColorBrightYellow, passMsg)

	totalPoints := correct*CORRECT_POINT + pass*PASS_POINT + wrong*WRONG_POINT

	totalPointsMsg := fmt.Sprintf("  Total Points: %d", totalPoints)
	colorize(ColorBrightCyan, totalPointsMsg)
}

func readCSV(fileName *string) *os.File {
	pwd, _ := os.Getwd()
	filePath := filepath.Join(pwd, *fileName)

	file, err := os.Open(filePath)
	if err != nil {
		msg := fmt.Sprintf("Failed to open the CSV file: %s!", filePath)
		colorize(ColorRed, msg)
		os.Exit(1)
	}

	return file
}

func parseCSV(file *os.File) []QuizRecord {
	quizs := []QuizRecord{}
	r := csv.NewReader(file)
	count := 1
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		Q, A := record[0], record[1]

		quizs = append(quizs, QuizRecord{
			ID: count,
			Q:  Q,
			A:  A,
		})

		count += 1
	}

	return quizs
}
