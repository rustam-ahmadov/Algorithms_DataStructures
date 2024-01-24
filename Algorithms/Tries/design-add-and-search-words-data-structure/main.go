package main

import "fmt"

// Constraints:

// 1 <= word.length <= 25
// word in addWord consists of lowercase English letters.
// word in search consist of '.' or lowercase English letters.
// There will be at most 2 dots in word for search queries.
// At most 104 calls will be made to addWord and search.

type WordDictionary struct {
	letter   byte
	childs   []WordDictionary
	terminal bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		childs: make([]WordDictionary, 0),
	}
}

func (this *WordDictionary) AddWord(word string) {
	if len(word) == 0 {
		this.terminal = true
		return
	}

	n := len(this.childs)
	for i := 0; i < n; i++ {
		if this.childs[i].letter == word[0] {
			this.childs[i].AddWord(word[1:])
			return
		}
	}
	this.childs = append(this.childs, WordDictionary{letter: word[0]})
	this.AddWord(word)
}

func (this *WordDictionary) Search(word string) bool {
	if len(word) == 0 {
		return this.terminal
	}

	n := len(this.childs)
	if word[0] == '.' {
		for i := 0; i < n; i++ {
			if this.childs[i].Search(word[1:]) {
				return true
			}
		}
		return false
	}

	for i := 0; i < n; i++ {
		if this.childs[i].letter == word[0] {
			return this.childs[i].Search(word[1:])
		}
	}
	return false
}

func main() {
	wordDictionary := Constructor()
	wordDictionary.AddWord("bad")
	wordDictionary.AddWord("dad")
	wordDictionary.AddWord("mad")
	b := wordDictionary.Search("pad") // return False
	fmt.Println(b)
	b = wordDictionary.Search("bad")  // return True
	fmt.Println(b)
	b = wordDictionary.Search(".ad")  // return True
	fmt.Println(b)
	b = wordDictionary.Search("b..")  // return True
	fmt.Println(b)


}
