package main

import (
	"fmt"
	"sort"
)

// There are n cars going to the same destination along a one-lane road. The destination is target miles away.
// You are given two integer array position and speed, both of length n, where position[i] is
// the position of the ith car and speed[i] is the speed of the ith car (in miles per hour).
// A car can never pass another car ahead of it, but it can catch up to it and drive bumper to bumper at the same speed.
// The faster car will slow down to match the slower car's speed. The distance between these two cars is ignored (i.e., they are assumed to have the same position).
// A car fleet is some non-empty set of cars driving at the same position and same speed. Note that a single car is also a car fleet.
// If a car catches up to a car fleet right at the destination point, it will still be considered as one car fleet.
// Return the number of car fleets that will arrive at the destination.

//Constraints:
// 	n == position.length == speed.length
// 	1 <= n <= 105
// 	0 < target <= 106
// 	0 <= position[i] < target
// 	All the values of position are unique.
// 	0 < speed[i] <= 106

func carFleet(target int, position []int, speed []int) int {
	var n int = len(position)
	if n == 1 {
		return 1
	}
	//pos - speed
	var m map[int]int = make(map[int]int)
	for i := 0; i < n; i++ {
		m[position[i]] = speed[i]
	}
	//sort desc, we need to start from last position
	sort.Slice(position, func(i, j int) bool {
		return position[i] > position[j]
	})

	//find the steps first car(closest to the target)need to go till the target
	var nextCarSteps float32 = (float32(target-position[0]) / float32(m[position[0]]))
	
	numFleets := n
	for i := 1; i < n; i++ {
		curStep := float32(target-position[i]) / float32(m[position[i]])
		if curStep <= nextCarSteps {
			numFleets--
		}else {
			nextCarSteps = curStep
		}
	}
	return numFleets
}

func main() {
	speed := []int{2,1,3}
	position := []int{0,4,2}
	target := 10

	fmt.Println(carFleet(target, position, speed))
}
