package boggle

type GamePlayer struct {
	Name  string
	Words []string
}

type GameOptions struct {
	GameLength          uint8
	MinWordLength       uint8
	ExpectedPlayerCount uint8
}

type Game struct {
	Id     string
	Matrix [16]string
	// maybe another radix tree of words found
	// constructed via solver?
	Words   []string
	Players map[string]GamePlayer
	Options GameOptions
}
