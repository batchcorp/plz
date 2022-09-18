package main

import (
	"github.com/batchcorp/plz/witch"
	"os"
)

func main() {
	addr := os.Getenv("WITCH_VIEWER")
	if addr == "" {
		addr = ":8318"
	}
	witch.StartViewer(addr)
}
