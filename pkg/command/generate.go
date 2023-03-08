package command

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/webhippie/prom-to-apt-dater/pkg/output"
	"github.com/webhippie/prom-to-apt-dater/pkg/prometheus"
)

var (
	generateCmd = &cobra.Command{
		Use:     "generate",
		Aliases: []string{"gen"},
		Short:   "Generate a host config",
		Run:     generateAction,
		Args:    cobra.NoArgs,
	}

	defaultGeneratePrometheusURL      = "http://localhost:9090"
	defaultGeneratePrometheusUsername = ""
	defaultGeneratePrometheusPassword = ""
	defaultGenerateOutputFilter       = "1 == 1"
	defaultGenerateOutputGroup        = "job"
	defaultGenerateOutputName         = "[__address__]"
	defaultGenerateOutputUser         = "static('root')"
	defaultGenerateOutputHost         = "[__address__]"
	defaultGenerateOutputPort         = "static('22')"
	defaultGenerateOutputFile         = ""
	defaultGenerateOutputTemplate     = ""
)

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.PersistentFlags().String("prometheus-url", defaultGeneratePrometheusURL, "URL to access Prometheus")
	viper.SetDefault("prometheus.url", defaultGeneratePrometheusURL)
	viper.BindPFlag("prometheus.url", generateCmd.PersistentFlags().Lookup("prometheus-url"))

	generateCmd.PersistentFlags().String("prometheus-username", defaultGeneratePrometheusUsername, "Username to access Prometheus")
	viper.SetDefault("prometheus.username", defaultGeneratePrometheusUsername)
	viper.BindPFlag("prometheus.username", generateCmd.PersistentFlags().Lookup("prometheus-username"))

	generateCmd.PersistentFlags().String("prometheus-password", defaultGeneratePrometheusPassword, "Password to access Prometheus")
	viper.SetDefault("prometheus.password", defaultGeneratePrometheusPassword)
	viper.BindPFlag("prometheus.password", generateCmd.PersistentFlags().Lookup("prometheus-password"))

	generateCmd.PersistentFlags().String("output-filter", defaultGenerateOutputFilter, "Query to filter host results")
	viper.SetDefault("output.filter", defaultGenerateOutputFilter)
	viper.BindPFlag("output.filter", generateCmd.PersistentFlags().Lookup("output-filter"))

	generateCmd.PersistentFlags().String("output-group", defaultGenerateOutputGroup, "Attribute to define the group")
	viper.SetDefault("output.group", defaultGenerateOutputGroup)
	viper.BindPFlag("output.group", generateCmd.PersistentFlags().Lookup("output-group"))

	generateCmd.PersistentFlags().String("output-name", defaultGenerateOutputName, "Attribute to name the host")
	viper.SetDefault("output.name", defaultGenerateOutputName)
	viper.BindPFlag("output.name", generateCmd.PersistentFlags().Lookup("output-name"))

	generateCmd.PersistentFlags().String("output-user", defaultGenerateOutputUser, "Attribute to detect the user")
	viper.SetDefault("output.user", defaultGenerateOutputUser)
	viper.BindPFlag("output.user", generateCmd.PersistentFlags().Lookup("output-user"))

	generateCmd.PersistentFlags().String("output-host", defaultGenerateOutputHost, "Attribute to access the host")
	viper.SetDefault("output.host", defaultGenerateOutputHost)
	viper.BindPFlag("output.host", generateCmd.PersistentFlags().Lookup("output-host"))

	generateCmd.PersistentFlags().String("output-port", defaultGenerateOutputPort, "Attribute to detect the port")
	viper.SetDefault("output.port", defaultGenerateOutputPort)
	viper.BindPFlag("output.port", generateCmd.PersistentFlags().Lookup("output-port"))

	generateCmd.PersistentFlags().String("output-file", defaultGenerateOutputFile, "Path to generated hosts file")
	viper.SetDefault("output.file", defaultGenerateOutputFile)
	viper.BindPFlag("output.file", generateCmd.PersistentFlags().Lookup("output-file"))

	generateCmd.PersistentFlags().String("output-template", defaultGenerateOutputTemplate, "Path to optional hosts template")
	viper.SetDefault("output.template", defaultGenerateOutputTemplate)
	viper.BindPFlag("output.template", generateCmd.PersistentFlags().Lookup("output-template"))
}

func generateAction(_ *cobra.Command, _ []string) {
	prom, err := prometheus.New(
		prometheus.WithURL(
			viper.GetString("prometheus.url"),
		),
		prometheus.WithUsername(
			viper.GetString("prometheus.username"),
		),
		prometheus.WithPassword(
			viper.GetString("prometheus.password"),
		),
	)

	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to initialize prometheus client")

		os.Exit(1)
	}

	targets, err := prom.Targets()

	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to fetch targets from prometheus")

		os.Exit(1)
	}

	outp, err := output.New(
		output.WithFilter(
			viper.GetString("output.filter"),
		),
		output.WithGroup(
			viper.GetString("output.group"),
		),
		output.WithName(
			viper.GetString("output.name"),
		),
		output.WithUser(
			viper.GetString("output.user"),
		),
		output.WithHost(
			viper.GetString("output.host"),
		),
		output.WithPort(
			viper.GetString("output.port"),
		),
		output.WithFile(
			viper.GetString("output.file"),
		),
		output.WithTemplate(
			viper.GetString("output.template"),
		),
		output.WithTargets(
			targets,
		),
	)

	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to initialize output client")

		os.Exit(1)
	}

	if err := outp.Generate(); err != nil {
		log.Error().
			Err(err).
			Msg("failed to generate hosts configuration")

		os.Exit(1)
	}
}
