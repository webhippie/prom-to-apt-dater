package config

// Output defines the output stuff like filters.
type Output struct {
	Filter   string `mapstructure:"filter"`
	Group    string `mapstructure:"group"`
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	File     string `mapstructure:"file"`
	Template string `mapstructure:"template"`
}

// Prometheus defines the prometheus connection details.
type Prometheus struct {
	URL      string `mapstructure:"url"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

// Logs defines the level and color for log configuration.
type Logs struct {
	Level  string `mapstructure:"level"`
	Pretty bool   `mapstructure:"pretty"`
	Color  bool   `mapstructure:"color"`
}

// Config is a combination of all available configurations.
type Config struct {
	Output     Output     `mapstructure:"output"`
	Prometheus Prometheus `mapstructure:"prometheus"`
	Logs       Logs       `mapstructure:"log"`
}

// Load initializes a default configuration struct.
func Load() *Config {
	return &Config{}
}
