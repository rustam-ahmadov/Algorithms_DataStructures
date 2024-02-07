package main

import (
	"fmt"
	"sort"
)

// Constraints:

// 1 <= intervals.length <= 104
// intervals[i].length == 2
// 0 <= starti <= endi <= 104

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{intervals[0]}
	rI := 0
	for i := 1; i < len(intervals); i++ {
		//if cur in the interval
		if intervals[i][0] >= res[rI][0] && intervals[i][1] <= res[rI][1] {
			continue
		}else if intervals[i][0] <= res[rI][1] {
			res[rI][1] = intervals[i][1]
		} else {
			res = append(res, intervals[i])
			rI++
		}
	}
	return res
}

func main() {
	// intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	intervals := [][]int{{1, 4}, {2, 3}}
	fmt.Println(merge(intervals))
}
