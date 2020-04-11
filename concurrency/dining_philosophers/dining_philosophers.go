package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	num             int
	perm            chan bool
	leftCS, rightCS *ChopS
}

var wg sync.WaitGroup

const amount = 5
const concurrentlyEating = 2

/*
5 philosophers and chopsticks naturally limit the number of concurrently eating to 2 anyway.
So in order to see the host limit in action lift the overall amount from 5 to 6 (or more for
a higher and quicker chance) and then compare results when concurrentlyEating is 2 and 3
*/

func (p Philo) eat() {
	for i := 0; i < 3; i++ {
		if rand.Intn(2) == 0 {
			p.leftCS.Lock()
			p.rightCS.Lock()
		} else {
			p.rightCS.Lock()
			p.leftCS.Lock()
		}

		p.perm <- false
		fmt.Println("starting to eat ", p.num)
		time.Sleep(time.Millisecond * 10)
		fmt.Println("finishing eating", p.num)
		<-p.perm

		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
	wg.Done()
}

func main() {
	CSticks := make([]*ChopS, amount)
	for i := 0; i < amount; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, amount)
	perm := make(chan bool, concurrentlyEating)
	go host(perm)
	for i := 0; i < amount; i++ {
		philos[i] = &Philo{
			i + 1,
			perm,
			CSticks[i],
			CSticks[(i+1)%amount],
		}
	}
	wg.Add(amount)
	for i := 0; i < amount; i++ {
		go philos[i].eat()
	}
	wg.Wait()
}

func host(perm chan bool) {
	for {
		perm <- true
		<-perm
	}
}
