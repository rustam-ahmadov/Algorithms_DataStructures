/*
The classical Dining philosophers problem.
*/

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const (
	philosophersCount = 5
	totalTime         = time.Second * 10
	slowMotionTime    = time.Millisecond * 20
	timeToEat         = time.Millisecond * 500
)

type action string

var (
	thinking action = "thinking"
	eating   action = "eating"
)

type Fork struct {
	sync.Mutex
	id int
}

type RisePlate struct {
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

func (p *Philosopher) eat() {
	//right
	p.rightFork.Lock()
	defer p.rightFork.Unlock()
	fmt.Println(p.name, " has taken right fork with id: ", p.rightFork.id)

	time.Sleep(time.Millisecond * 10)

	//left
	p.leftFork.Lock()
	defer p.leftFork.Unlock()
	fmt.Println(p.name, " has taken left fork with id: ", p.leftFork.id)

	//eat
	fmt.Println(p.name, " is eating", "...")
	time.Sleep(timeToEat)
}

func (p *Philosopher) think() {
	fmt.Println(p.name, " is thinking", "...")
}

// region version with deadlock
func eatWithDeadlock(ctx context.Context, p *Philosopher, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		return
	default:
		p.think()
		p.eat()
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

func eatWithSemaphore(ctx context.Context, ch chan struct{}, p *Philosopher, wg *sync.WaitGroup, totalRise *int) {
	defer wg.Done()

	// if we need to do some work
	for p.plate.grams > 0 {

		select {
		case <-ctx.Done():
			return
		default:
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
}

// startMealSession with deadlock
func startMealSessionWithSemaphore(ctx context.Context, ps [5]Philosopher) {
	var totalRise int
	for i := range len(ps) {
		totalRise += int(ps[i].plate.grams)
	}

	wg := &sync.WaitGroup{}
	wg.Add(philosophersCount)

	//semaphore
	ch := make(chan struct{}, 4)

	// give freedom and life to philosophers :)
	go eatWithSemaphore(ctx, ch, &ps[0], wg, &totalRise)
	go eatWithSemaphore(ctx, ch, &ps[1], wg, &totalRise)
	go eatWithSemaphore(ctx, ch, &ps[2], wg, &totalRise)
	go eatWithSemaphore(ctx, ch, &ps[3], wg, &totalRise)
	go eatWithSemaphore(ctx, ch, &ps[4], wg, &totalRise)

	wg.Wait()
	close(ch)

	fmt.Println("totalRise: ", totalRise)
}

//endregion

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), totalTime)
	defer cancel()

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

	start := time.Now()
	startMealSessionWithSemaphore(ctx, philosophers)
	fmt.Println(time.Since(start))
}
