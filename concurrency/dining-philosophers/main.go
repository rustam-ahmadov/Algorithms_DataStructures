/*
The classical Dining philosophers problem.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	philosophersCount = 5
	slowMotionTime    = time.Millisecond * 20
)

type action string
type grams int

var (
	thinking action = "thinking"
	eating   action = "eating"
)

type Fork struct {
	capacity grams
	sync.Mutex
	id int
}

func newFork(id int, capacity grams) Fork {
	return Fork{
		capacity: capacity,
		id:       id,
	}
}

type RisePlate struct {
	grams grams
}

func newRisePlate(gr grams) RisePlate {
	return RisePlate{grams: gr}
}

type Philosopher struct {
	name                string
	plate               *RisePlate
	rightFork, leftFork *Fork
}

func newPhilosopher(name string, risePlate *RisePlate, rightFork, leftFork *Fork) Philosopher {
	return Philosopher{
		name:      name,
		plate:     risePlate,
		rightFork: rightFork,
		leftFork:  leftFork,
	}
}

var mu sync.Mutex

// region version with deadlock

func eatWithDeadlock(p *Philosopher, wg *sync.WaitGroup, totalRise *int) {
	defer wg.Done()

	// if we need to do some work
	for p.plate.grams > 0 {

		//thinking
		fmt.Println(p.name, " is thinking", "...")

		//trying to take the forks
		p.rightFork.Lock()
		time.Sleep(time.Millisecond * 10)
		fmt.Println(p.name, " has taken right fork with id: ", p.rightFork.id)
		p.leftFork.Lock()
		fmt.Println(p.name, " has taken left fork with id: ", p.leftFork.id)
		//----

		//starting to eat
		fmt.Println(p.name, " is eating", "...")

		time.Sleep(slowMotionTime)

		totalCap := p.leftFork.capacity + p.rightFork.capacity

		mu.Lock()
		*totalRise -= int(totalCap)
		mu.Unlock()
		p.plate.grams -= totalCap

		//free the forks
		p.leftFork.Unlock()
		time.Sleep(time.Millisecond * 10)
		p.rightFork.Unlock()
	}
}

// startMealSessionDeadlock with deadlock
func startMealSessionDeadlock(ps [5]Philosopher) {
	var totalRise int
	for i := range len(ps) {
		totalRise += int(ps[i].plate.grams)
	}

	wg := &sync.WaitGroup{}
	wg.Add(philosophersCount)

	// give freedom and life to philosophers :)
	go eatWithDeadlock(&ps[0], wg, &totalRise)
	go eatWithDeadlock(&ps[1], wg, &totalRise)
	go eatWithDeadlock(&ps[2], wg, &totalRise)
	go eatWithDeadlock(&ps[3], wg, &totalRise)
	go eatWithDeadlock(&ps[4], wg, &totalRise)

	wg.Wait()

	fmt.Println("totalRise: ", totalRise)
}

//endregion

// region version with Semaphore

func eatWithSemaphore(ch chan struct{}, p *Philosopher, wg *sync.WaitGroup, totalRise *int) {
	defer wg.Done()

	// if we need to do some work
	for p.plate.grams > 0 {

		ch <- struct{}{}

		//thinking
		fmt.Println(p.name, " is thinking", "...")

		//trying to take the forks
		p.rightFork.Lock()
		time.Sleep(time.Millisecond * 10)
		fmt.Println(p.name, " has taken right fork with id: ", p.rightFork.id)
		p.leftFork.Lock()
		fmt.Println(p.name, " has taken left fork with id: ", p.leftFork.id)
		//----

		//starting to eat
		fmt.Println(p.name, " is eating", "...")

		time.Sleep(slowMotionTime)

		totalCap := p.leftFork.capacity + p.rightFork.capacity

		mu.Lock()
		*totalRise -= int(totalCap)
		mu.Unlock()

		p.plate.grams -= totalCap

		fmt.Println(*totalRise)

		//free the forks
		p.leftFork.Unlock()
		time.Sleep(time.Millisecond * 10)
		p.rightFork.Unlock()

		<-ch
	}
}

// startMealSession with deadlock
func startMealSessionWithSemaphore(ps [5]Philosopher) {
	var totalRise int
	for i := range len(ps) {
		totalRise += int(ps[i].plate.grams)
	}

	wg := &sync.WaitGroup{}
	wg.Add(philosophersCount)

	//semaphore
	ch := make(chan struct{}, 2)

	// give freedom and life to philosophers :)
	go eatWithSemaphore(ch, &ps[0], wg, &totalRise)
	go eatWithSemaphore(ch, &ps[1], wg, &totalRise)
	go eatWithSemaphore(ch, &ps[2], wg, &totalRise)
	go eatWithSemaphore(ch, &ps[3], wg, &totalRise)
	go eatWithSemaphore(ch, &ps[4], wg, &totalRise)

	wg.Wait()
	close(ch)

	fmt.Println("totalRise: ", totalRise)
}

//endregion

func main() {

	fork1 := newFork(0, 5)
	fork2 := newFork(1, 5)
	fork3 := newFork(2, 5)
	fork4 := newFork(3, 5)
	fork5 := newFork(4, 5)

	risePlate1 := newRisePlate(1000)
	risePlate2 := newRisePlate(500)
	risePlate3 := newRisePlate(500)
	risePlate4 := newRisePlate(600)
	risePlate5 := newRisePlate(200)

	p1 := newPhilosopher("0_Tural", &risePlate1, &fork1, &fork2)
	p2 := newPhilosopher("1_Eldar", &risePlate2, &fork2, &fork3)
	p3 := newPhilosopher("2_Rustam", &risePlate3, &fork3, &fork4)
	p4 := newPhilosopher("3_Kamil", &risePlate4, &fork4, &fork5)
	p5 := newPhilosopher("4_Ulvi", &risePlate5, &fork5, &fork1)

	philosophers := [5]Philosopher{p1, p2, p3, p4, p5}

	startMealSessionWithSemaphore(philosophers)
}
