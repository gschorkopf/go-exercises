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

	wordsWithCount := map[string]int{}
	for _, word := range words {
		wordsWithCount[word]++
	}

	for word, count := range wordsWithCount {
		fmt.Println(word, "has a count of", count)
	}

}
