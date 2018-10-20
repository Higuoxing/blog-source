package main

import (
	"fmt"
	"runtime"
	"time"
	"math/rand"
)

type semaphore chan int

var (
	smoker_match = make(semaphore, 1)
	smoker_paper = make(semaphore, 1)
	smoker_tobacco = make(semaphore, 1)
	smoking_done = make(semaphore, 1)
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

func provider() {
	for {
		random := rand.Intn(3)
		switch (random) {
		case 0:
			smoker_match.V()
		case 1:
			smoker_paper.V()
		case 2:
			smoker_tobacco.V()
		}
		smoking_done.P()
	}
}

func smoker_0() {
	for {
		smoker_match.P()
		fmt.Println("Smoker who has match is smoking")
		smoking_done.V()
	}
}

func smoker_1() {
	for {
		smoker_paper.P()
		fmt.Println("Smoker who has paper is smoking")
		smoking_done.V()
	}
}

func smoker_2() {
	for {
		smoker_tobacco.P()
		fmt.Println("Smoker who has tobacco is smoking")
		smoking_done.V()
	}
}

func init() {
	smoking_done.V()
}

func main() {
	go provider()
	go smoker_0()
	go smoker_1()
	go smoker_2()
	time.Sleep(time.Duration(5) * time.Millisecond)
	return
}
