/*!
 * main.go
 * binarySearchTree
 *
 * Copyright (c) 2019 Masaharu TASHIRO
 * Released under the MIT license.
 * see https://opensource.org/licenses/MIT
 */
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
	left  *Node
	right *Node
	value int
}

func main() {
	// Parse arguments
	if len(os.Args) < 3 {
		log.Println("Too few arguments")
		return
	} else if len(os.Args) > 3 {
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

	target, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Println(err)
		return
	}
	if target < 1 {
		fmt.Println("TargetValue must be greater than 0")
		return
	}

	// Initilize list
	tree := (*Node)(nil)
	rand.Seed(time.Now().UnixNano())
	tree.insertRandomValues(n)

	// Print unsorted list
	fmt.Println("Tree")
	tree.print()

	// searching
	searchedNode := tree.search(5)

	// Print search result
	fmt.Println("Search result")
	searchedNode.print()

}

func (tree *Node) insert(value int) {
	if tree == nil {
		fmt.Println("Founded nil")
		tree = &Node{
			value: value,
			left:  nil,
			right: nil,
		}
		tree.print()
	} else {
		if value < tree.value {
			fmt.Println("left")
			tree.left.insert(value)
		} else {
			fmt.Println("right")
			tree.right.insert(value)
		}
	}
}

func (tree *Node) insertRandomValues(n int) {
	for i := 0; i < n; i++ {
		tree.insert(rand.Intn(2 * n))
	}
}

func (tree *Node) print() {
	if tree == nil {
		//fmt.Println("Reach to nil")
	} else {
		fmt.Print(tree.value)
		tree.left.print()
		tree.right.print()
	}
}

func (tree *Node) search(value int) *Node {
	searchedNode := tree
	return searchedNode
}
