package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Quiz Game Start!")

	// fileName := "problems.csv"
	fileName := flag.String("csv", "./problems/main.csv", "the input csv file, the format is question,answer.")
	timeLimit := flag.Int("limit", 180, "the question time limit")
	flag.Parse()

	if *timeLimit < 1 {
		colorize(ColorRed, "The time should be positive")
		return
	}

	file := readCSV(fileName)
	quizs := parseCSV(file)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(quizs), func(i, j int) { quizs[i], quizs[j] = quizs[j], quizs[i] })

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
			if s != q.A {
				reply := fmt.Sprintf("The answer is: %s", q.A)
				colorize(ColorRed, "Wrong!")
				colorize(ColorBrightMagneta, reply)
				wrong++
				continue
			}

			colorize(ColorGreen, "Correct!")
			correct++
		case <-time.After(time.Duration(*timeLimit) * time.Second):
			fmt.Println("")
			colorize(ColorRed, "Time out!")
			PrintResult(len(quizs), correct, 0, len(quizs)-correct)
			return
		}
	}

	PrintResult(len(quizs), correct, pass, wrong)

}
