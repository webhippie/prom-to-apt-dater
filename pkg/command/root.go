package command

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/owncloud/ocis-hello/pkg/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Root is the entry point for the ocis-hello command.
func Root() *cobra.Command {
	if env := os.Getenv("PROM_TO_APTDATER_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	cmd := &cobra.Command{
		Use:     "prom-to-apt-dater",
		Short:   "Prometheus targets to APT-Dater hosts",
		Long:    ``,
		Version: version.String,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			setupLogger()
			setupConfig()
		},
	}

	cmd.Flags().BoolP("version", "v", false, "Print the current version")
	cmd.PersistentFlags().BoolP("help", "h", false, "Show the help output")

	cmd.PersistentFlags().String("log-level", "", "Set logging level")
	viper.BindPFlag("log.level", cmd.PersistentFlags().Lookup("log-level"))
	viper.SetDefault("log.level", "info")
	viper.BindEnv("log.level", "PROM_TO_APTDATER_LOG_LEVEL")

	cmd.PersistentFlags().Bool("log-pretty", false, "Enable pretty logging")
	viper.BindPFlag("log.pretty", cmd.PersistentFlags().Lookup("log-pretty"))
	viper.SetDefault("log.pretty", true)
	viper.BindEnv("log.pretty", "PROM_TO_APTDATER_LOG_PRETTY")

	cmd.PersistentFlags().Bool("log-color", false, "Enable colored logging")
	viper.BindPFlag("log.color", cmd.PersistentFlags().Lookup("log-color"))
	viper.SetDefault("log.color", true)
	viper.BindEnv("log.color", "PROM_TO_APTDATER_LOG_COLOR")

	cmd.AddCommand(Generate())

	return cmd
}
