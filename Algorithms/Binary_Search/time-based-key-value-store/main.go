package main

import "fmt"

type TimeMap struct {
	m map[string][]Node
}
type Node struct {
	value     string
	timestamp int
}

func Constructor() TimeMap {
	return TimeMap{
		m: make(map[string][]Node),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	nodes := this.m[key]
	nodes = append(nodes, Node{value, timestamp})
	this.m[key] = nodes

}

func (this *TimeMap) Get(key string, timestamp int) string {
	nodes, isPresent := this.m[key]
	if !isPresent {
		return ""
	}

	l, r := 0, len(nodes)-1
	if nodes[r].timestamp <= timestamp {
		return nodes[r].value
	}

	res := ""
	for l <= r {
		m := (l + r) / 2

		if timestamp == nodes[m].timestamp {
			return nodes[m].value
		}

		if nodes[m].timestamp > timestamp {
			r = m - 1
		} else if nodes[m].timestamp < timestamp {
			res = nodes[m].value
			l = m + 1
		}
	}
	return res
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */

func main() {

	// [[],["foo","bar",1],["foo",1],["foo",3],["foo","bar2",4],["foo",4],["foo",5]]
	timeMap := Constructor()
	timeMap.Set("love", "high", 10)
	timeMap.Set("love", "low", 20)
	val := timeMap.Get("love", 5)
	fmt.Print(val)

}
