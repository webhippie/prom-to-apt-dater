package main

import (
	"os"

	"github.com/webhippie/prom-to-apt-dater/pkg/command"
)

func main() {
	if err := command.Root().Execute(); err != nil {
		os.Exit(1)
	}
}
