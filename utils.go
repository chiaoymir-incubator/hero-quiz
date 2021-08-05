package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// func colorize(color Color, message string) {
// 	fmt.Printf("%s%s%s", string(color), message, string(ColorReset))
// }

func colorizeln(color Color, message string) {
	fmt.Printf("%s%s%s\n", string(color), message, string(ColorReset))
}

func readCSV(fileName *string) *os.File {
	pwd, _ := os.Getwd()
	filePath := filepath.Join(pwd, *fileName)

	file, err := os.Open(filePath)
	if err != nil {
		msg := fmt.Sprintf("Failed to open the CSV file: %s!", filePath)
		colorizeln(ColorRed, msg)
		os.Exit(1)
	}

	return file
}

func answerNotCorrect(ans string, reply string) bool {
	return ans != reply
}
