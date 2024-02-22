package main

import "fmt"

type trienode struct {
	Chieldren []*trienode  
	Terminal bool
}

func main() {
	node := trienode{}
	fmt.Print(node.Chieldren)
}
