package main

import "fmt"

type trie struct {
	root *trieNode
}

type trieNode struct {
	children map[string]*trieNode
}

/*
search algo:
	- set currentNode to equal the node that called the function
	- iterate over the word's chars
	- see if currentnode contains the char
	- if it does, set currentNode to equal the child at that char
	- if it does not, return false
	- after iterating, see if currentNode contains "*". If it does, return true. Otherwise, return false
*/
func (t *trieNode) search(word string) bool {
	currentNode := t
	for _, val := range word {
		if _, ok := currentNode.children[string(val)]; ok {
			currentNode = currentNode.children[string(val)]
			continue
		}
		return false
	}
	if _, ok := currentNode.children["*"]; ok {
		return true
	}

	return false
}

func (t *trieNode) insert(word string) {

	if len(word) == 0 {
		t.children["*"] = &trieNode{}
		return
	}

	firstLetter := string([]rune(word)[0])

	if _, ok := t.children[firstLetter]; !ok {
		t.children[firstLetter] = &trieNode{make(map[string]*trieNode)}
	}

	t.children[firstLetter].insert(string([]rune(word)[1:]))
}

func (t *trieNode) collectWords(word string, collection *[]string) {

	if _, ok := t.children["*"]; ok {
		*collection = append(*collection, word)
	}
	for k := range t.children {
		t.children[k].collectWords(word+k, collection)
	}
}

func main() {
	var myTrie trie
	myTrie.root = &trieNode{make(map[string]*trieNode)}

	myTrie.root.insert("bad")
	myTrie.root.insert("let")
	myTrie.root.insert("heat")
	myTrie.root.insert("myself")
	words := []string{}
	myTrie.root.collectWords("", &words)
	fmt.Println(words)
	fmt.Println(myTrie.root.search("heat"))
}
