package config

type StoreConfig struct {
	DatabaseURL string `toml:"database_url"`
}

func NewStoreConfig() *StoreConfig {
	return &StoreConfig{
		DatabaseURL: "postgres://user:password@localhost:5432/dbname",
	}
}