package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	file, err := os.Open("1K_English.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	allWords := []string{}

	for scanner.Scan() {
		word := scanner.Text()
		if word != "" {
			allWords = append(allWords, word)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	rand.NewSource((time.Now().UnixNano()))
	rand.Shuffle(len(allWords), func(i, j int) {
		allWords[i], allWords[j] = allWords[j], allWords[i]
	})

	wordsToType := allWords[:25]

	fmt.Println(wordsToType)
	fmt.Println("enter words:")

	var input string
	startTime := time.Now()

	fmt.Scanln(&input)

	duration := time.Since(startTime).Minutes()
	wpm := float64(25) / duration

	fmt.Println("Stats:")
	fmt.Printf("Time: %.2fs\n", time.Since(startTime).Seconds())
	fmt.Printf("WPM: %.2f\n", wpm)
}
