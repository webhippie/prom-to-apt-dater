package command

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/webhippie/prom-to-apt-dater/pkg/output"
	"github.com/webhippie/prom-to-apt-dater/pkg/prometheus"
)

// Generate is the entrypoint for the generate command.
func Generate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate a host config",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			prom, err := prometheus.New(
				prometheus.WithURL(viper.GetString("prometheus.url")),
				prometheus.WithUsername(viper.GetString("prometheus.username")),
				prometheus.WithPassword(viper.GetString("prometheus.password")),
			)

			if err != nil {
				log.Error().
					Err(err).
					Msg("Failed to initialize prometheus client")

				return err
			}

			targets, err := prom.Targets()

			if err != nil {
				log.Error().
					Err(err).
					Msg("Failed to fetch targets from prometheus")

				return err
			}

			outp, err := output.New(
				output.WithFilter(viper.GetString("output.filter")),
				output.WithGroup(viper.GetString("output.group")),
				output.WithName(viper.GetString("output.name")),
				output.WithHost(viper.GetString("output.host")),
				output.WithFile(viper.GetString("output.file")),
				output.WithTemplate(viper.GetString("output.template")),
				output.WithTargets(targets),
			)

			if err != nil {
				log.Error().
					Err(err).
					Msg("Failed to initialize output client")

				return err
			}

			if err := outp.Generate(); err != nil {
				log.Error().
					Err(err).
					Msg("Failed to generate hosts configuration")

				return err
			}

			return nil
		},
	}

	cmd.Flags().String("prometheus-url", "", "URL to access Prometheus")
	viper.BindPFlag("prometheus.url", cmd.Flags().Lookup("prometheus-url"))
	viper.BindEnv("prometheus.url", "PROM_TO_APTDATER_PROMETHEUS_URL")

	cmd.Flags().String("prometheus-username", "", "Username to access Prometheus")
	viper.BindPFlag("prometheus.username", cmd.Flags().Lookup("prometheus-username"))
	viper.BindEnv("prometheus.username", "PROM_TO_APTDATER_PROMETHEUS_USERNAME")

	cmd.Flags().String("prometheus-password", "", "Password to access Prometheus")
	viper.BindPFlag("prometheus.password", cmd.Flags().Lookup("prometheus-password"))
	viper.BindEnv("prometheus.password", "PROM_TO_APTDATER_PROMETHEUS_PASSWORD")

	cmd.Flags().String("output-filter", "", "Query to filter host results")
	viper.BindPFlag("output.filter", cmd.Flags().Lookup("output-filter"))
	viper.BindEnv("output.filter", "PROM_TO_APTDATER_OUTPUT_FILTER")

	cmd.Flags().String("output-group", "", "Attribute to define the group")
	viper.BindPFlag("output.group", cmd.Flags().Lookup("output-group"))
	viper.BindEnv("output.group", "PROM_TO_APTDATER_OUTPUT_GROUP")

	cmd.Flags().String("output-name", "", "Attribute to name the host")
	viper.BindPFlag("output.name", cmd.Flags().Lookup("output-name"))
	viper.BindEnv("output.name", "PROM_TO_APTDATER_OUTPUT_NAME")

	cmd.Flags().String("output-host", "", "Attribute to access the host")
	viper.BindPFlag("output.host", cmd.Flags().Lookup("output-host"))
	viper.BindEnv("output.host", "PROM_TO_APTDATER_OUTPUT_HOST")

	cmd.Flags().String("output-file", "", "Path to generated hosts file")
	viper.BindPFlag("output.file", cmd.Flags().Lookup("output-file"))
	viper.BindEnv("output.file", "PROM_TO_APTDATER_OUTPUT_FILE")

	cmd.Flags().String("output-template", "", "Path to optional hosts template")
	viper.BindPFlag("output.template", cmd.Flags().Lookup("output-template"))
	viper.BindEnv("output.template", "PROM_TO_APTDATER_OUTPUT_TEMPLATE")

	return cmd
}

// init defined the default options for viper.
func init() {
	viper.SetDefault("prometheus.url", "http://localhost:9090")
	viper.SetDefault("prometheus.username", "")
	viper.SetDefault("prometheus.password", "")

	viper.SetDefault("output.filter", "")
	viper.SetDefault("output.group", "job")
	viper.SetDefault("output.name", "[__address__]")
	viper.SetDefault("output.host", "[__address__]")
	viper.SetDefault("output.file", "")
	viper.SetDefault("output.template", "")
}
