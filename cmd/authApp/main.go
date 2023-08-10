package main

import (
	"auth/internal/authApp/server"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	log.Fatal(server.Start())
}
