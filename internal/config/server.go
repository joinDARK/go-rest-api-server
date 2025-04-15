package config

type ServerConfig struct {
	Addr     string `toml:"addr"`
	LogLevel string `toml:"log_level"`
	Store *StoreConfig
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		Addr:     ":8000",
		LogLevel: "debug",
		Store: NewStoreConfig(),
	}
}

