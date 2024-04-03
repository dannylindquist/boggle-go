package game

import (
	"fmt"
	"math/rand"
)

var Dice = [16][6]string{
	{"R", "I", "F", "O", "B", "X"},
	{"I", "F", "E", "H", "E", "Y"},
	{"D", "E", "N", "O", "W", "S"},
	{"U", "T", "O", "K", "N", "D"},
	{"H", "M", "S", "R", "A", "O"},
	{"L", "U", "P", "E", "T", "S"},
	{"A", "C", "I", "T", "O", "A"},
	{"Y", "L", "G", "K", "U", "E"},
	{"Qu", "B", "M", "J", "O", "A"},
	{"E", "H", "I", "S", "P", "N"},
	{"V", "E", "T", "I", "G", "N"},
	{"B", "A", "L", "I", "Y", "T"},
	{"E", "Z", "A", "V", "N", "D"},
	{"R", "A", "L", "E", "S", "C"},
	{"U", "W", "I", "L", "R", "G"},
	{"P", "A", "C", "E", "M", "D"},
}

func GenerateMatrix() [16]string {
	// new letters
	matrix := [16]string{}
	// track what index in the matrix has been filled
	usedIndex := make(map[int]bool)
	// place each die
	fmt.Println("starting filling matrix")
	for num := range len(Dice) {
		// random number between 0 and len(Dice)
		diceIndex := rand.Intn(len(Dice) - num)
		// if diceIndex is already used, increment up
		loops := 0
		for usedIndex[diceIndex] {
			if diceIndex == len(Dice) {
				diceIndex = 0
			} else {
				diceIndex = diceIndex + 1
			}
			loops = loops + 1
			if loops > 100 {
				fmt.Println("we stuck dude")
			}
		}
		side := rand.Intn(6)
		matrix[num] = Dice[diceIndex][side]
		usedIndex[diceIndex] = true
	}
	return matrix
}
