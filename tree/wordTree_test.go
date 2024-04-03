package tree

import (
	"os"
	"testing"
)

func Test_wordTreeInit(t *testing.T) {
	file, err := os.OpenFile("../words.txt", os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	tree := NewWordTree(file)
	if tree.head.Children == nil {
		t.Errorf("head.children should not be nil")
	}
}

func Test_wordTreeFindWord(t *testing.T) {
	file, err := os.OpenFile("../words.txt", os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	tree := NewWordTree(file)
	if !tree.StartsWith("hell") {
		t.Errorf("StartsWith should return false")
	}
}
