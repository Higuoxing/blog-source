package main

import (
	"fmt"
	"runtime"
	"time"
)

type semaphore chan int

var (
	chopstics = make([]semaphore, 5)
	mutex   = make(semaphore, 1)
)

func (sema *semaphore) P() {
	for {
		if len(*sema) > 0 {
			<- *sema
			break
		}
		runtime.Gosched()
	}
}

func (sema *semaphore) V() {
	for {
		if len(*sema) < cap(*sema) {
			*sema <- 1
			break
		}
		runtime.Gosched()
	}
}

func dining(i int) {
	for {
		mutex.P()
		chopstics[i].P()
		chopstics[(i+1)%5].P()
		mutex.V()
		fmt.Printf("Philosopher %v is eating\n", i+1)
		chopstics[i].V()
		chopstics[(i+1)%5].V()
		fmt.Printf("Philosopher %v is thinking\n", i+1)
	}
}

func init() {
	for i := 0; i < 5; i ++ {
		chopstics[i] = make(semaphore, 1)
		chopstics[i].V()
	}
	mutex.V()
}

func main() {
	for i := 0; i < 5; i ++ {
		go dining(i)
	}
	time.Sleep(time.Duration(5) * time.Millisecond)
	return
}
