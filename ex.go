package main

import (
	"bufio"
	"fmt"
	"os"
)

func countwords() {
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
func countlines() {
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
func countchar() {
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
func vowels(char int32) bool {

	if (char == 'a') || (char == 'e') || (char == 'i') ||
		(char == 'o') || (char == 'u') {

		return true

	} else {

		return false

	}

}
func punc(p int32) bool {

	if (p == '!') || (p == '.') || (p == '-') ||
		(p == ',') || (p == ';') || (p == ':') {

		return true

	} else {

		return false

	}

}

// type allvowels struct {
// 	v1, v2, v3, v4, v5 int32
// }

func main() {
	Content, _ := os.ReadFile("text.txt")
	s := string(Content)
	// fmt.Println(s)
	// No of Words
	countwords()
	// No of lines
	countlines()
	// No of characters
	countchar()
	// No of Vowels
	// v := allvowels{'a', 'e', 'i', 'o', 'u'}
	countvowels := 0
	for _, char := range s {
		if vowels(char) {
			countvowels++
		}
	}

	fmt.Println("No of vowels :", countvowels)
	// No of Puntuation
	countpunc := 0
	for _, p := range s {
		if punc(p) {
			countpunc++
		}
	}
	fmt.Println("No of puntuations:", countpunc)
}
