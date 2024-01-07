package main

import "container/list"

type Node struct {
	key int
	val int
}

type LRUCache struct {
	capacity int
	m        map[int]*list.Element
	list     *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		m:        make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (lru *LRUCache) Get(key int) int {
	el, ok := lru.m[key]
	if ok {
		lru.list.MoveToBack(el)
		return el.Value.(Node).val
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	el, ok := lru.m[key]
	if ok { //update
		el.Value = Node{key: key, val: value}
		lru.list.MoveToBack(el)
		return
	}

	if lru.list.Len() == lru.capacity {
		first := lru.list.Front()
		delete(lru.m, first.Value.(Node).key)
		lru.list.Remove(first)
	}
	// put
	lru.m[key] = lru.list.PushBack(Node{key: key, val: value})
}

func main() {

}
