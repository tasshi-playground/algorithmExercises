package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Println("argument should be <text> <pattern>")
		return
	}
	text := ([]rune)(os.Args[1])
	pattern := ([]rune)(os.Args[2])

	if len(text) < len(pattern) {
		log.Println("<text> should be longer than <pattern>")
		return
	}

	skipMap := createSkipMap(pattern)
	fmt.Println("skipMap")
	for key, val := range skipMap {
		fmt.Printf("%c: %d\n", key, val)
	}

}

func createSkipMap(pattern []rune) map[rune]int {
	skipMap := make(map[rune]int, len(pattern))
	for index, word := range pattern {
		if index != len(pattern)-1 {
			skipMap[word] = len(pattern) - index - 1
		}
	}
	return skipMap
}
