package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dannylindquist/boggle-go/server"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT)

	fmt.Println("starting server")
	s := server.NewServer(8080)
	go s.ListenAndServe()

	<-sig
}
