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
	"errors"
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
	n, target, err := parseArgs(os.Args)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("N: %d, Target: %d\n", n, target)

	// Initilize list
	tree := (*Node)(nil)
	rand.Seed(time.Now().UnixNano())
	insertRandomValues(&tree, n)

	// Print unsorted list
	fmt.Println("Tree")
	tree.print()
	fmt.Println()

	// searching
	searchedNode := tree.search(target)

	// Print search result
	fmt.Println("Search result")
	if searchedNode == nil {
		fmt.Println("Not found")
	} else {
		searchedNode.print()
	}
	fmt.Println()

}

func insert(tree **Node, value int) {
	if *tree == (*Node)(nil) {
		//fmt.Println("Founded nil")
		*tree = &Node{
			value: value,
			left:  nil,
			right: nil,
		}
		//(*tree).print()
	} else {
		if value < (*tree).value {
			//fmt.Println("left")
			insert(&(*tree).left, value)
		} else {
			//fmt.Println("right")
			insert(&(*tree).right, value)
		}
	}
}

func insertRandomValues(tree **Node, n int) {
	for i := 0; i < n; i++ {
		insert(tree, rand.Intn(2*n))
	}
}

func (tree *Node) print() {
	if tree == (*Node)(nil) {
		fmt.Print("(nil)")
	} else {
		fmt.Print(tree.value)
		tree.left.print()
		tree.right.print()
	}
}

func (tree *Node) search(value int) *Node {
	if tree == (*Node)(nil) {
		return nil
	}
	switch {
	case tree.value == value:
		return tree
	case tree.value > value:
		return tree.left.search(value)
	case tree.value < value:
		return tree.right.search(value)
	}
	return nil
}

func parseArgs(args []string) (int, int, error) {

	if len(args) < 3 {
		err := errors.New("Too few arguments")
		return 0, 0, err
	} else if len(args) > 3 {
		err := errors.New("Too much arguments")
		return 0, 0, err
	}
	n, err := strconv.Atoi(args[1])
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}
	if n < 1 {
		err := errors.New("N must be greater than 0")
		return 0, 0, err
	}

	target, err := strconv.Atoi(args[2])
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}
	if target < 1 || target >= n {
		err := errors.New("TargetValue must be greater than 0 and smaller than N")
		return 0, 0, err
	}
	return n, target, nil
}
