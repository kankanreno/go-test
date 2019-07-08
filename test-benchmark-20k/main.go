package main

import (
	"bufio"
	"log"
	"os"
	"runtime/pprof"
)

type Node struct {
	Char rune
	Children []*Node
}

func NewNode(r rune) *Node {
	return &Node{Char: r}
}

func (n *Node) insert(r rune) *Node {
	child := n.get(r)
	if child == nil {
		child = NewNode(r)
		n.Children = append(n.Children, child)
	}

	return child
}

func (n *Node) get(r rune) *Node {
	for _, child := range n.Children {
		if child.Char == r {
			return child
		}
	}

	return nil
}

type Trie struct {
	Root *Node
}

func NewTrie() *Trie {
	var r rune
	trie := Trie{Root: NewNode(r)}
	return &trie
}

func (tr *Trie) Build(word string) {
	node := tr.Root
	runeArr := []rune(word)
	for _, char := range runeArr {
		child := node.insert(char)
		node = child
	}
}

func (tr *Trie) Has(word string) bool {
	node := tr.Root
	runeArr := []rune(word)
	for _, char := range runeArr {
		found := node.get(char)
		if found == nil {
			return false
		}
		node = found
	}
	return true
}

func main() {
	cpuProfile, _ := os.Create("cpu_profile")
	pprof.StartCPUProfile(cpuProfile)
	defer pprof.StopCPUProfile()

	var trie1 = NewTrie()
	file, err := os.Open("./20k.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		trie1.Build(scanner.Text())
	}

	trie1.Has("42082")
	trie1.Has("oops")
	trie1.Has("Supercalifragilisticexpialidocious")
}









































