// Copyright 2019 Masaharu TASHIRO
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	sortedList := mergeSort(list)

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
			// fmt.Print("left :")
			// printList(list)
			// fmt.Print("right:")
			// printList(list2)
			return list, list2
		}
		middle = middle.next
		end = end.next
		if end != nil {
			end = end.next
		}
	}
}

func mergeSort(list *Node) *Node {
	left, right := divide(list)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	left = mergeSort(left)
	right = mergeSort(right)
	// fmt.Print("left :")
	// printList(left)
	// fmt.Print("right:")
	// printList(right)

	mergedList := &Node{}
	end := mergedList
	for {
		if left == nil {
			end.next = right
			// fmt.Print("merge:")
			// printList(mergedList.next)
			return mergedList.next
		}
		if right == nil {
			end.next = left
			// fmt.Print("merge:")
			// printList(mergedList.next)
			return mergedList.next
		}
		if left.value < right.value {
			end.next = left
			end = end.next
			left = left.next
		} else {
			end.next = right
			end = end.next
			right = right.next
		}
	}
}

func printList(list *Node) {
	for node := list; node != nil; node = node.next {
		fmt.Print(node.value, ",")
	}
	fmt.Println()
}
