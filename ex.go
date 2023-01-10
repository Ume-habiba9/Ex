package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func count(file string, wg *sync.WaitGroup) (int, int, int, int) {
	countWords := 0
	countVowel := 0
	countPunc := 0
	countLines := 0
	for _, char := range file {
		if (char == 'a') || (char == 'e') || (char == 'i') ||
			(char == 'o') || (char == 'u') {
			countVowel++

		}
		if (char == '!') || (char == '.') || (char == '-') ||
			(char == ',') || (char == ';') || (char == ':') {
			countPunc++
		}
		if char == ' ' {
			countWords++
		}
		if char == '\n' {
			countLines++
		}
	}
	v, _ := fmt.Println("No of Vowels:", countVowel)
	p, _ := fmt.Println("No of Puntuation:", countPunc)
	w, _ := fmt.Println("No of Words:", countWords)
	l, _ := fmt.Println("No of LInes:", countLines)
	defer wg.Done()

	return v, p, w, l
}

func main() {
	startTime := time.Now()
	var wg sync.WaitGroup
	fileContent, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Println(err)
	}
	file := string(fileContent)
	wg.Add(1)
	count(file, &wg)
	fmt.Println("Number of characters:", len(file))
	wg.Wait()
	ExecutionTime := time.Since(startTime)
	fmt.Println("Execution Time of Whole Program is :", ExecutionTime)

}
