package command

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// setupLogger prepares the logger.
func setupLogger() {
	switch strings.ToLower(viper.GetString("log.level")) {
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	if viper.GetBool("log.pretty") {
		log.Logger = log.Output(
			zerolog.ConsoleWriter{
				Out:     os.Stderr,
				NoColor: !viper.GetBool("log.color"),
			},
		)
	}
}

// setupConfig prepares the config.
func setupConfig() {
	viper.SetConfigName("config")

	viper.AddConfigPath("/etc/prom-to-apt-dater")
	viper.AddConfigPath("$HOME/.prom-to-apt-dater")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			log.Debug().
				Msg("Continue without config file")
		case viper.UnsupportedConfigError:
			log.Fatal().
				Msg("Unsupported config type")
		default:
			if e := log.Debug(); e.Enabled() {
				log.Fatal().
					Err(err).
					Msg("Failed to read config")
			} else {
				log.Fatal().
					Msg("Failed to read config")
			}
		}
	}
}
