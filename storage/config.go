package storage

type Config struct {
	//Строка подключения к базе данных
	DatabaseURI string `toml:"database_uri"`
}

func NewConfig() *Config {
	return &Config{}
}
