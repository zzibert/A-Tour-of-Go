package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	leftCS, rightCS *ChopS
	index           int
}

func (p Philo) eat() {
	fmt.Println("starting to eat ", p.index)
	for i := 0; i < 3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("eating ", p.index)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
	wg.Done()
	fmt.Println("finished eating ", p.index)
}

func host(i int, guard chan int, philos []*Philo) {
	wg.Add(1)
	go philos[i].eat()
	wg.Wait()
	<-guard
}

func main() {
	guard := make(chan int, 2)
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5], i}
	}
	for i := 0; i < 5; i++ {
		guard <- i
		host(i, guard, philos)
	}
}
