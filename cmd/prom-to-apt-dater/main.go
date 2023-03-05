package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/webhippie/prom-to-apt-dater/pkg/command"
)

func main() {
	if env := os.Getenv("PROM_TO_APT_DATER_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	if err := command.Run(); err != nil {
		os.Exit(1)
	}
}
