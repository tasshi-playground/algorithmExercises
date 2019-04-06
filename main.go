package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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

func initList(int) *Node {
	return nil
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
