package server

import (
	"net/http"

	"github.com/dannylindquist/boggle-go/store"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer(port int) *Server {

	mux := http.NewServeMux()

	gameStore := store.NewGameStore()
	registerRoutes(mux, gameStore)
	server := &Server{
		mux: mux,
	}

	return server
}

func (s *Server) ListenAndServe() {
	err := http.ListenAndServe(":3333", s.mux)
	if err != nil {
		panic(err)
	}
}
