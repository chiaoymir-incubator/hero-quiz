package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func parseCSVQuiz(file *os.File) []QuizRecord {
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

func printResult(totalQ, correct, pass, wrong int) {
	fmt.Println("")
	fmt.Println("Result:")

	totalQMsg := fmt.Sprintf("  Number of Questions: %d", totalQ)
	colorizeln(ColorBrightWhite, totalQMsg)

	correctMsg := fmt.Sprintf("  Correct: %d", correct)
	wrongMsg := fmt.Sprintf("  Wrong: %d", wrong)
	passMsg := fmt.Sprintf("  Pass: %d", pass)
	colorizeln(ColorBrightGreen, correctMsg)
	colorizeln(ColorBrightRed, wrongMsg)
	colorizeln(ColorBrightYellow, passMsg)

	totalPoints := correct*CORRECT_POINT + pass*PASS_POINT + wrong*WRONG_POINT

	totalPointsMsg := fmt.Sprintf("  Total Points: %d", totalPoints)
	colorizeln(ColorBrightCyan, totalPointsMsg)
}

func startQuiz(quizs []QuizRecord, t int, n int) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(quizs), func(i, j int) { quizs[i], quizs[j] = quizs[j], quizs[i] })

	if n <= len(quizs) {
		quizs = quizs[0:n]
	}

	scanner := bufio.NewScanner(os.Stdin)
	correct := 0
	wrong := 0
	pass := 0
	for i, q := range quizs {
		fmt.Printf("Problem #%d, @%d: %s\n = ", i+1, q.ID, q.Q)
		ch := make(chan string, 1)

		go func() {
			if scanner.Scan() {
				s := strings.TrimSpace(scanner.Text())
				ch <- s
			}
		}()

		select {
		case s := <-ch:
			if answerNotCorrect(s, q.A) {
				reply := fmt.Sprintf("The answer is: %s", q.A)
				colorizeln(ColorRed, "Wrong!")
				colorizeln(ColorBrightMagneta, reply)
				wrong++
				continue
			}

			colorizeln(ColorGreen, "Correct!")
			correct++
		case <-time.After(time.Duration(t) * time.Second):
			fmt.Println("")
			colorizeln(ColorRed, "Time out!")
			printResult(len(quizs), correct, 0, len(quizs)-correct)
			return
		}
	}

	printResult(len(quizs), correct, pass, wrong)
}
