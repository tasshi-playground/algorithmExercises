/*!
 * main.go
 * sleepSortByLinkedList
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
	"sync"
	"time"
)

// Node ...
type Node struct {
	value int
	next  *Node
}

func main() {
	// Parse arguments
	n, err := parseArgs(os.Args)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("N: %d\n", n)

	// Initilize list
	rand.Seed(time.Now().UnixNano())
	list := initList(n)

	// Print unsorted list
	fmt.Println("Unsorted list")
	printList(list)

	// sorting
	sortedList := sleepSort(list)

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

func sleepSort(list *Node) *Node {
	sortedList := &Node{}
	end := sortedList

	// semaphores
	var wwgg sync.WaitGroup
	var wg sync.WaitGroup
	mutex := &sync.RWMutex{}
	locker := &sync.Mutex{}
	cond := sync.NewCond(locker)

	wwgg.Add(1)
	go func() {
		defer wwgg.Done()
		for node := list; node != nil; node = node.next {
			wg.Add(1)
			wwgg.Add(1)
			go func(value int) {
				locker.Lock()
				defer wg.Done()
				wwgg.Done()
				cond.Wait()
				locker.Unlock()

				time.Sleep((time.Duration)(value) * time.Millisecond)

				mutex.Lock()
				end.next = &Node{value: value, next: nil}
				end = end.next
				mutex.Unlock()
				//fmt.Println(value)
			}(node.value)
		}
	}()

	//fmt.Println("Wait for make goroutine")
	wwgg.Wait()
	//fmt.Println("start goroutine")
	cond.Broadcast()
	//fmt.Println("Wait for finish goroutine")
	wg.Wait()

	return sortedList.next
}

func printList(list *Node) {
	for node := list; node != nil; node = node.next {
		fmt.Print(node.value, ",")
	}
	fmt.Println()
}

func parseArgs(args []string) (int, error) {

	if len(args) < 2 {
		err := errors.New("Too few arguments")
		return 0, err
	} else if len(args) > 2 {
		err := errors.New("Too much arguments")
		return 0, err
	}
	n, err := strconv.Atoi(args[1])
	if err != nil {
		log.Println(err)
		return 0, err
	}
	if n < 1 {
		err := errors.New("N must be greater than 0")
		return 0, err
	}

	return n, nil
}
