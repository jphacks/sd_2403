package main

import (
	"log"
	"sushi/server"
)

func main() {
	log.Fatal(server.Start())
}
