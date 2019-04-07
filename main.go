package main

import (
	"fmt"
	"sync"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	leftCS, rightCS *ChopS
}

func (p Philo) eat() {
	for {
		p.leftCS.lock()
		p.rightCS.lock()

		fmt.Println("eating")

		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
}

func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5]}
	}
	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}
}
