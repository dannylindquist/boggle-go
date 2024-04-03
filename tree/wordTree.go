package tree

import (
	"bufio"
	"io"
	"strings"
)

type Node struct {
	IsWord   bool
	IsCommon bool
	Children map[string]*Node
}

type WordTree struct {
	head *Node
}

func NewWordTree(reader io.Reader) *WordTree {
	tree := &WordTree{
		head: &Node{
			Children: make(map[string]*Node),
		},
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		word := scanner.Text()
		common := strings.HasSuffix(word, "+")
		word = strings.TrimRight(word, "+")
		tree.AddWord(word, common)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return tree
}

func (t *WordTree) AddWord(word string, isCommon bool) {
	current := t.head
	for _, letter := range word {
		if _, ok := current.Children[string(letter)]; !ok {
			current.Children[string(letter)] = &Node{
				Children: make(map[string]*Node),
			}
		}
		current = current.Children[string(letter)]
	}
	current.IsWord = true
	current.IsCommon = isCommon
}

func (t *WordTree) StartsWith(partial string) bool {
	node := t.head
	for _, letter := range partial {
		if _, ok := node.Children[string(letter)]; !ok {
			return false
		}
		node = node.Children[string(letter)]
	}
	return true
}

func (t *WordTree) FindWord(word string) *Node {
	current := t.head
	for _, letter := range word {
		if _, ok := current.Children[string(letter)]; !ok {
			return nil
		}
		current = current.Children[string(letter)]
	}
	return current
}
