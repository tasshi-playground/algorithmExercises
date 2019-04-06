package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Node ...
type Node struct {
	value int
	next  *Node
}

func main() {
	// Parse arguments
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
	if n < 1 {
		fmt.Println("N must be greater than 0")
		return
	}

	// Initilize list
	rand.Seed(time.Now().UnixNano())
	list := initList(n)

	// Print unsorted list
	fmt.Println("Unsorted list")
	for node := list; node != nil; node = node.next {
		fmt.Print(node.value, ",")
	}
	fmt.Println()

	// sorting
	sortedList, err := mergeSort(list)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print sorted list
	fmt.Println("Sorted list")
	for node := sortedList; node != nil; node = node.next {
		fmt.Print(node.value, ",")
	}
	fmt.Println()

}

func initList(n int) *Node {
	list := (*Node)(nil)
	for i := 0; i < n; i++ {
		list = &Node{value: rand.Intn(2 * n), next: list}
	}
	return list
}

func divide(*Node) (*Node, *Node) {
	return nil, nil
}

func mergeOrderByAsc(*Node, *Node) *Node {
	return nil
}

func mergeSort(*Node) (*Node, error) {
	return nil, nil
}
