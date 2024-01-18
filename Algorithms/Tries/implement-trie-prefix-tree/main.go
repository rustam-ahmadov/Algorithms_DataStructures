package main

import "fmt"

// Constraints:

// 1 <= word.length, prefix.length <= 2000
// word and prefix consist only of lowercase English letters.
// At most 3 * 104 calls in total will be made to insert, search, and startsWith.

type Trie struct {
	letter   byte //lowercase english letters only
	childs   []Trie
	terminal bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	n := len(word)
	if n == 0 {
		this.terminal = true
		return
	}
	for i := 0; i < len(this.childs) -1 ; i++ {
		if this.childs[i].letter == word[0] {
			this.childs[i].Insert(word[1:])
			return
		}
	}
	
	this.childs = append(this.childs, Trie{letter: word[0]})
	this.Insert(word)
}

func (this *Trie) Search(word string) bool {
	n := len(word)
	if n == 1 {
		for i := 0; i < len(this.childs) -1; i++ {
			if this.childs[i].letter == word[0]  {
				// fmt.Println(word[0], this.childs[i].letter)
				return this.childs[i].terminal == true
			}
		}
	}else {
		for i := 0; i < len(this.childs) -1; i++ {
			if this.childs[i].letter == word[0]{
				// fmt.Println(word[0], this.childs[i].letter)
				return this.childs[i].Search(word[1:])
			}
		}
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	n := len(prefix)
	if n == 0 {
		return true
	}
	for i := 0; i < len(this.childs) -1; i++ {
		if this.childs[i].letter == prefix[0]{
			return this.childs[i].StartsWith(prefix[1:])
		}
	}
	return false
}

func main() {
	trie := Constructor()
	fmt.Println("insert apple")
	trie.Insert("apple")
	fmt.Print("search apple: ")
	fmt.Println(trie.Search("apple"))
	fmt.Print("search app: ")
	fmt.Println(trie.Search("app"))
	fmt.Print("starts with app: ")
	fmt.Println(trie.StartsWith("app"))
	fmt.Println("insert app")
	trie.Insert("app")
	fmt.Print("search app: ")
	fmt.Println(trie.Search("app"))
	

	fmt.Println(trie.Search("a"))
	// var a byte = '1'
	// str:= "1gc"

	// fmt.Println(a == str[0])
}
