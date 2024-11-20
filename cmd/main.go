package main

import (
	"log"

	"github.com/damirpavlik/proxy-pilot/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalf("skill issue starting the server from main: %v", err)
	}
}
