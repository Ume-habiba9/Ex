package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func countWords(totalWords string) int {
	words := strings.Fields(totalWords)
	return len(words)
}

func countLines(totalLines string) int {
	lines := strings.Split(totalLines, "\n")
	return len(lines)
}
func countVowels(char int32) bool {

	if (char == 'a') || (char == 'e') || (char == 'i') ||
		(char == 'o') || (char == 'u') {

		return true

	} else {

		return false

	}

}
func countPunc(p int32) bool {

	if (p == '!') || (p == '.') || (p == '-') ||
		(p == ',') || (p == ';') || (p == ':') {

		return true

	} else {

		return false

	}

}

func main() {
	startTime := time.Now()
	fileContent, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Println(err)
	}
	file := string(fileContent)
	// No of Words
	fmt.Println("Number of words:", countWords(file))
	// No of lines
	fmt.Println("Number of Lines:", countLines(file))
	// No of characters
	fmt.Println("Number of characters:", len(file))
	// No of Vowels
	countvowels := 0
	for _, char := range file {
		if countVowels(char) {
			countvowels++
		}
	}
	fmt.Println("No of vowels :", countvowels)
	// No of Puntuation
	countpunc := 0
	for _, punctuation := range file {
		if countPunc(punctuation) {
			countpunc++
		}
	}
	fmt.Println("No of puntuations:", countpunc)
	ExecutionTime := time.Since(startTime)
	fmt.Println("Execution Time of Whole Program is :", ExecutionTime)
}
