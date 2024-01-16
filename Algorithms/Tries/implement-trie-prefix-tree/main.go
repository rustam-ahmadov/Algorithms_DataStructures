package main

import "fmt"

// Constraints:

// 1 <= word.length, prefix.length <= 2000
// word and prefix consist only of lowercase English letters.
// At most 3 * 104 calls in total will be made to insert, search, and startsWith.

type node struct {
	letter   byte //lowercase english letters only
	childs   []node
	terminal bool
}

type Trie struct {
	letters []node
}

func Constructor() Trie {
	return Trie{
		letters: make([]node, 0),
	}
}

func (this *Trie) Insert(word string) {
	n := len(word)
	if n == 0 {
		return
	}

	for i := 0; i < len(this.letters); i++ {
		if this.letters[i].letter == word[0] {
			this.Insert(word[1:])
			return
		}
	}
	this.letters = append(this.letters, node{letter: word[0]})
	this.Insert(word)
}

// func (this *Trie) Search(word string) bool {

// }

// func (this *Trie) StartsWith(prefix string) bool {

// }

func main() {
	trie := Constructor()
	trie.Insert("abc")
	fmt.Println(trie)

}
