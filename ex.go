package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

func countWords(totalWords string, wg *sync.WaitGroup) int {
	words := strings.Fields(totalWords)
	w, _ := fmt.Println("Number of words:", len(words))
	wg.Done()
	return w
}

func countLines(totalLines string, wg *sync.WaitGroup) int {
	lines := strings.Split(totalLines, "\n")
	l, _ := fmt.Println("Number of Lines:", len(lines))
	wg.Done()
	return l
}
func countVowelsandPunc(file string, wg *sync.WaitGroup) (int, int) {
	countVowel := 0
	countPunc := 0
	for _, char := range file {
		if (char == 'a') || (char == 'e') || (char == 'i') ||
			(char == 'o') || (char == 'u') {
			countVowel++

		}
	}
	v, _ := fmt.Println("No of Vowels:", countVowel)
	for _, p := range file {
		if (p == '!') || (p == '.') || (p == '-') ||
			(p == ',') || (p == ';') || (p == ':') {
			countPunc++
		}
	}
	p, _ := fmt.Println("No of Puntuation:", countPunc)
	wg.Done()
	return v, p
}

func main() {
	startTime := time.Now()
	var wg sync.WaitGroup
	fileContent, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Println(err)
	}
	file := string(fileContent)
	// No of Words
	wg.Add(1)
	go countWords(file, &wg)
	wg.Wait()
	// No of lines
	wg.Add(1)
	go countLines(file, &wg)
	wg.Wait()
	// No of characters
	fmt.Println("Number of characters:", len(file))
	// No of Vowels
	wg.Add(1)
	go countVowelsandPunc(file, &wg)
	wg.Wait()
	// No of Puntuations
	go countVowelsandPunc(file, &wg)
	ExecutionTime := time.Since(startTime)
	fmt.Println("Execution Time of Whole Program is :", ExecutionTime)
	wg.Wait()
}
