package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Too few arguments")
		return
	} else if len(os.Args) > 2 {
		log.Println("Too much arguments")
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Println(err)
		return
	}
	if n < 3 {
		fmt.Println("N must be greater than or equal to 3")
		return
	}

	magicSquare, err := resolveMagicSquare(n)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the Result
	fmt.Printf("%dx%d Magic Square\n", n, n)
	for _, ms := range magicSquare {
		fmt.Println(ms)
	}

}

func resolveMagicSquare(n int) ([][]int, error) {
	if n < 3 {
		return nil, fmt.Errorf("Error: %s", "N must be greater than or equal to 3")
	}
	if n%2 == 0 {
		return nil, fmt.Errorf("Error: %s", "N must be odd number")
	}
	// Initialize matrix
	var magicSquare = make([][]int, n)
	for i := range magicSquare {
		magicSquare[i] = make([]int, n)
	}

	maxNumber := n * n
	posX := (n+1)/2 - 1
	posY := n - 1

	for num := 1; num <= maxNumber; num++ {
		magicSquare[posY][posX] = num
		nextPosX := (posX + 1) % n
		nextPosY := (posY + 1) % n
		if magicSquare[nextPosY][nextPosX] == 0 {
			posX = nextPosX
			posY = nextPosY
		} else {
			posY = int(math.Max(float64(posY-1), 0))
		}
		fmt.Println(posX)
		magicSquare[posY][posX] = num
	}

	return magicSquare, nil
}
