package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type chopStick struct{ sync.Mutex }

type philo struct {
	leftCS, rightCS *chopStick
	number          int
	needToEat       int
	wg              *sync.WaitGroup
	lock            *sync.Mutex
}

func (p *philo) eat() {
	p.lock.Lock()
	if p.needToEat == 0 {
		return
	}
	p.leftCS.Lock()
	p.rightCS.Lock()

	fmt.Printf("starting to eat %d - %d\n", p.number, p.needToEat)
	fmt.Printf("finishing eating %d - %d\n", p.number, p.needToEat)

	p.rightCS.Unlock()
	p.leftCS.Unlock()
	p.needToEat = p.needToEat - 1
	p.wg.Done()
	p.lock.Unlock()
}

func main() {
	//Create chopsticks
	cSticks := make([]*chopStick, 5)
	for i := 0; i < 5; i++ {
		cSticks[i] = new(chopStick)
	}

	//Create philosophers
	philos := make([]*philo, 5)
	for i := 0; i < 5; i++ {
		w := &sync.WaitGroup{}
		w.Add(3)
		philos[i] = &philo{cSticks[i], cSticks[(i+1)%5], i, 3, w, &sync.Mutex{}}
	}

	go host(philos, cSticks)
	for _, v := range philos {
		v.wg.Wait()
	}
}

func host(philos []*philo, cSticks []*chopStick) {
	for {
		randomPhilo := rand.Intn(len(philos))
		secondPhilo := rand.Intn(len(philos))
		go philos[randomPhilo].eat()
		go philos[secondPhilo].eat()
	}
}
