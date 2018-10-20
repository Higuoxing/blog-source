package main

import (
	"time"
	"fmt"
)

type semaphore chan int

const maxSize = 3

var (
	emptyCount = make(semaphore, maxSize)
	fullCount  = make(semaphore, maxSize)
	useQueue   = make(semaphore, 1)
	items      = make(semaphore, maxSize)
)

func (sema *semaphore) P() {
	for {
		if len(*sema) > 0 {
			<- *sema
			break
		}
	}
}

func (sema *semaphore) V() {
	for {
		if len(*sema) < cap(*sema) {
			*sema <- 1
			break
		}
	}
}

func producer() {
	for {
		emptyCount.P()
		useQueue.P()
		items.V()
		fmt.Printf("[Producer] Produce 1. Now we have %v items.\n", len(items))
		useQueue.V()
		fullCount.V()
	}
}

func consumer() {
	for {
		fullCount.P()
		useQueue.P()
		items.P()
		fmt.Printf("[Consumer] Consume 1. Now we have %v items.\n", len(items))
		useQueue.V()
		emptyCount.V()
	}
}

func init() {
	for i := 0; i < maxSize; i ++ {
		emptyCount <- 1
	}
	useQueue <- 1
}

func main() {
	go producer()
	go consumer()
	return
}
