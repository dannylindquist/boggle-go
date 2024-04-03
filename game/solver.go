package game

import (
	"math"
	"strconv"
	"strings"

	"github.com/dannylindquist/boggle-go/tree"
)

var directions = [8][2]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

type FoundWord struct {
	Word     string
	isCommon bool
}

type PossibleWord struct {
	Word string
	Path string
}

type Solver struct {
	wordTree   *tree.WordTree
	WordsFound []FoundWord
	Path       string
	matrix     [16]string
}

func NewSolver(wordTree *tree.WordTree, matrix [16]string) Solver {
	return Solver{
		wordTree:   wordTree,
		WordsFound: make([]FoundWord, 0),
		matrix:     matrix,
	}
}

func (s *Solver) Solve() {
	uniqueWords := make(map[string]bool)
	for i := range 16 {
		possible := make([]PossibleWord, 0, 10)
		s.walkIndex("", i, &possible, 0, "")
		for _, word := range possible {
			isWord := s.wordTree.FindWord(word.Word)
			if isWord.IsWord && !uniqueWords[word.Word] {
				s.WordsFound = append(s.WordsFound,
					FoundWord{
						Word:     word.Word,
						isCommon: isWord.IsCommon,
					})
				uniqueWords[word.Word] = true
			}
		}
	}
}

func (s Solver) walkIndex(current string, index int, possibleWords *[]PossibleWord, hitIndicies uint16, path string) {
	if (hitIndicies&(1<<index)) > 0 || index < 0 || index >= 16 {
		return
	}
	nextStep := strings.ToLower((current + s.matrix[index]))
	newPath := path + strconv.FormatInt(int64(index), 16)
	if len(nextStep) >= 3 {
		if !s.wordTree.StartsWith(nextStep) {
			return
		}

		*possibleWords = append(*possibleWords, PossibleWord{
			Word: nextStep,
			Path: newPath,
		})
	}
	visited := hitIndicies | (1 << index)

	for _, value := range directions {
		dx := value[0]
		dy := value[1]
		newRow := int(math.Floor(float64(index)/float64(4))) + dy
		newCol := (index % 4) + dx
		if newCol < 0 || newCol > 3 || newRow < 0 || newRow > 3 {
			continue
		}
		nextIndex := 4*newRow + newCol
		s.walkIndex(nextStep, nextIndex, possibleWords, visited, newPath)
	}
}
