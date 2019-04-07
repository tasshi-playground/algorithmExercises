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
	// fmt.Println("skipMap")
	// for key, val := range skipMap {
	// 	fmt.Printf("%c: %d\n", key, val)
	// }

	for i := 0; len(text)-len(pattern)-i >= 0; {
		skip := 0
		for index, word := range pattern {
			//fmt.Printf("%c, %c\n", word, text[i+index])
			if word != text[i+index] {
				if val, ok := skipMap[text[i+len(pattern)-1]]; ok {
					skip = val
				} else {
					skip = len(pattern)
				}
				break
			}
		}
		//fmt.Printf("skip: %d\n", skip)
		if skip == 0 {
			fmt.Printf("Match at: %d\n", i)
			return
		}
		i = i + skip
	}
	fmt.Println("No match")
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
