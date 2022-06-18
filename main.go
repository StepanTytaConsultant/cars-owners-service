package main

import (
	"os"

	"github.com/cars-owners-service/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(2)
	}
	os.Exit(0)
}
