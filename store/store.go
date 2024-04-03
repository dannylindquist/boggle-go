package store

import (
	"strconv"
	"sync"
	"time"

	"github.com/dannylindquist/boggle-go"
	"github.com/dannylindquist/boggle-go/game"
)

type GameStore struct {
	m         sync.Mutex
	currentId uint32
	Games     map[string]*boggle.Game
}

func NewGameStore() *GameStore {
	return &GameStore{
		currentId: 1,
		Games:     make(map[string]*boggle.Game),
	}
}

func (s *GameStore) AddGame(game *boggle.Game) {
	s.Games[game.Id] = game
}

func (s *GameStore) CreateGame(options boggle.GameOptions) *boggle.Game {
	s.m.Lock()
	defer s.m.Unlock()
	newId := s.currentId
	s.currentId = s.currentId + 1

	// generates random id based on time
	// includes incrementing id to avoid collisions
	timestamp := time.Now().UnixMilli()
	id := strconv.FormatUint(uint64(timestamp), 36) + strconv.Itoa(int(newId))
	newGame := &boggle.Game{
		Id:      id,
		Players: make(map[string]boggle.GamePlayer),
		Options: options,
	}
	newGame.Matrix = game.GenerateMatrix()

	s.Games[newGame.Id] = newGame
	return newGame
}

func (s *GameStore) GetGame(id string) *boggle.Game {
	if game, ok := s.Games[id]; ok {
		return game
	}
	return nil
}
