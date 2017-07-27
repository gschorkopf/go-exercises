package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter words: ")

	input, _ := reader.ReadString('\n')
	words := strings.Fields(input)

	var wordsWithCount = map[string]int{}
	for _, word := range words {
		found, ok := wordsWithCount[word]

		if ok {
			wordsWithCount[word] = found + 1
		} else {
			wordsWithCount[word] = 1
		}
	}

	for word, count := range wordsWithCount {
		fmt.Println(word, "has a count of", count)
	}

}
