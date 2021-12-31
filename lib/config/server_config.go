package config

var (
	// DefaultEnv - Default Environment Variable
	DefaultEnv = "dev"

	// DefaultHost - Default Host Variable
	DefaultHost = "localhost"

	// DefaultPort - Default Port number to run the server
	DefaultPort = 3000

	// DefaultAPIPath - Default API Path
	DefaultAPIPath = "/"
)

// Config - Structure for Configuration
type Config struct {
	Env     string
	Host    string
	Port    int
	APIPath string
}

// NewConfig is the constructor function to "Config" structure
func NewConfig(env, host string, port int, apiPath string) *Config {
	return &Config{
		Env:     env,
		Host:    host,
		Port:    port,
		APIPath: apiPath,
	}
}

// DefaultConfig will provide default values to Config structure
func DefaultConfig() *Config {
	return &Config{
		Env:     DefaultEnv,
		Host:    DefaultHost,
		Port:    DefaultPort,
		APIPath: DefaultAPIPath,
	}
}
