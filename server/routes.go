package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/dannylindquist/boggle-go"
	"github.com/dannylindquist/boggle-go/server/assets"
	"github.com/dannylindquist/boggle-go/server/views"
	"github.com/dannylindquist/boggle-go/store"
)

func registerRoutes(mux *http.ServeMux, gameStore *store.GameStore) {
	mux.HandleFunc("GET /assets/", func(w http.ResponseWriter, r *http.Request) {
		slog.Info(
			"request file",
			"path", r.URL.Path,
			"method", r.Method,
			"url", r.URL,
		)
		http.StripPrefix("/assets/",
			http.FileServer(http.FS(assets.AssetFS)),
		).ServeHTTP(w, r)
	})

	mux.HandleFunc("GET /{$}", loggerMiddleware(homeHandler()))
	mux.HandleFunc("GET /single", loggerMiddleware(singleHandler()))
	mux.HandleFunc("POST /single", loggerMiddleware(postSingleHandler(gameStore)))

	mux.HandleFunc("GET /game/{gameId}", loggerMiddleware(gameHandler(gameStore)))
}

func homeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		views.Render(w, r, "index.html", nil)
	}
}

func singleHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		views.Render(w, r, "single.html", nil)
	}
}

func gameHandler(gameStore *store.GameStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		gameId := r.PathValue("gameId")
		game := gameStore.GetGame(gameId)
		// if it doesn't exist go home
		if game == nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		fmt.Println(game.Matrix)
		views.Render(w, r, "game.html", game)
	}
}

func postSingleHandler(gameStore *store.GameStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		timeLengthString := r.FormValue("length")
		wordLengthString := r.FormValue("wordlength")
		timeLength, err := strconv.Atoi(timeLengthString)
		if err != nil {
			fmt.Println(err)
			http.Redirect(w, r, "/single", http.StatusFound)
			return
		}
		wordLength, err := strconv.Atoi(wordLengthString)
		if err != nil {
			fmt.Println(err)
			http.Redirect(w, r, "/single", http.StatusFound)
			return
		}
		game := gameStore.CreateGame(boggle.GameOptions{
			GameLength:          uint8(timeLength),
			MinWordLength:       uint8(wordLength),
			ExpectedPlayerCount: 1,
		})
		http.Redirect(w, r, "/game/"+game.Id, http.StatusFound)
	}
}

func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("request", "method", r.Method, "url", r.URL)
		next(w, r)
	})
}
