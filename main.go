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
	printList(list)

	// sorting
	sortedList, err := mergeSort(list)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print sorted list
	fmt.Println("Sorted list")
	printList(sortedList)

}

func initList(n int) *Node {
	list := (*Node)(nil)
	for i := 0; i < n; i++ {
		list = &Node{value: rand.Intn(2 * n), next: list}
	}
	return list
}

func divide(list *Node) (*Node, *Node) {
	if list == nil {
		return nil, nil
	}
	middle := list
	end := list
	if list != nil {
		end = list.next
		if end != nil {
			end = end.next
		}
	} else {
		return list, nil
	}

	for {
		if end == nil {
			list2 := middle.next
			middle.next = nil
			printList(list)
			printList(list2)
			return list, list2
		}
		middle = middle.next
		end = end.next
		if end != nil {
			end = end.next
		}
	}
}

func mergeOrderByAsc(*Node, *Node) *Node {
	return nil
}

func mergeSort(*Node) (*Node, error) {
	return nil, nil
}

func printList(list *Node) {
	for node := list; node != nil; node = node.next {
		fmt.Print(node.value, ",")
	}
	fmt.Println()
}
