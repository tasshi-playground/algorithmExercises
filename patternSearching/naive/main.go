/*!
 * main.go
 * patternSearching/naive
 *
 * Copyright (c) 2019 Masaharu TASHIRO
 * Released under the MIT license.
 * see https://opensource.org/licenses/MIT
 */

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
	fmt.Println("Text:    ", (string)(text))
	fmt.Println("Pattern: ", (string)(pattern))

	if len(text) < len(pattern) {
		log.Println("<text> should be longer than <pattern>")
		return
	}

	for i := 0; len(text)-len(pattern)-i >= 0; i++ {
		isMatch := true
		for index, word := range pattern {
			if word != text[i+index] {
				isMatch = false
				break
			}
		}
		if isMatch {
			fmt.Printf("Match at: %d\n", i)
			return
		}
	}
	fmt.Println("No match")
}
