package main

import (
	"log"

	"github.com/rhnvrm/docmonk/pkg/server"
)

func main() {
	log.Fatal(server.Serve())
}
