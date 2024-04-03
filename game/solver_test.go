package game

import (
	"os"
	"testing"

	"github.com/dannylindquist/boggle-go/tree"
)

func TestSolver_FindWords(t *testing.T) {
	file, err := os.OpenFile("../words.txt", os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	words := tree.NewWordTree(file)
	board := [16]string{
		"E", "T", "A", "U", "E", "O", "E", "O", "X", "A", "S", "L", "C", "Y", "I", "P",
	}
	solver := NewSolver(words, board)
	solver.Solve()

	if len(solver.WordsFound) == 0 {
		t.Errorf("expected some words, got %d", len(solver.WordsFound))
	}
}
