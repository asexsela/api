package api

// Generl instance for API server of REST application
type Config struct {
	//Port
	BindAddr string `toml:"bind_addr"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}
