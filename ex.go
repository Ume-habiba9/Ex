package main

import (
	"bufio"
	"fmt"
	"os"
)

func countWords() {
	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	fmt.Println("No of Words:", count)
}
func countLines() {
	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanlines := bufio.NewScanner(file)
	scanlines.Split(bufio.ScanLines)
	var lines int
	for scanlines.Scan() {
		lines++
	}
	fmt.Println("No of lines:", lines)
}
func countChar() {
	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanchar := bufio.NewScanner(file)
	scanchar.Split(bufio.ScanRunes)
	char := 0
	for scanchar.Scan() {
		char++
	}
	fmt.Println("No of characters:", char)
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

// type Reader interface {
// 	readfile()
// }

func main() {
	fileContent, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Println(err)
	}
	file := string(fileContent)
	// No of Words
	countWords()
	// No of lines
	countLines()
	// No of characters
	countChar()
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
}
