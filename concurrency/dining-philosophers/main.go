/*
The classical Dining philosophers problem.

Implemented with forks (aka chopsticks) as mutexes.
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Fork struct {
	sync.Mutex
}

type Philosopher struct {
	id                  int
	leftFork, rightFork *Fork
}


func (p Philosopher) dine() {
	say("thinking", p.id)
	randomPause(1)

	say("hungry", p.id)
	p.leftFork.Lock()
	say("has took left fork", p.id)
	p.rightFork.Lock()
	say("has took right fork", p.id)
	
	say("eating", p.id)
	randomPause(2)
	
	p.rightFork.Unlock()
	say("has put right fork", p.id)
	p.leftFork.Unlock()
	say("has put left fork", p.id)

	p.dine()
}

func randomPause(max int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(max*1000)))
}

func say(action string, id int) {
	fmt.Printf("#%d is %s\n", id, action)
}

func main() {
	// How many philosophers and forks
	count := 5

	// Create forks
	forks := make([]*Fork, count)
	for i := 0; i < count; i++ {
		forks[i] = new(Fork)
	}

	// Create philospoher, assign them 2 forks and send them to the dining table
	philosophers := make([]*Philosopher, count)
	for i := 0; i < count; i++ {
		philosophers[i] = &Philosopher{
			id:        i + 1,
			leftFork:  forks[i],
			rightFork: forks[(i+1)%count], //last one 4 + 1 % count == 5 % 5 = 0
		}
		go philosophers[i].dine()
	}

	// Wait endlessly while they're dining
	endless := make(chan int)
	<-endless
}
