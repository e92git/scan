package apiserver

import "github.com/BurntSushi/toml"

// Config ...
type Config struct {
	BindAddr       string `toml:"bind_addr"`
	LogLevel       string `toml:"log_level"`
	RunCron        bool   `toml:"run_cron"`
	Dsn            string `toml:"dsn"`
	DsnTest        string `toml:"dsn_test"`
	ApiKeyAutocode string `toml:"api_key_autocode"`
	ApiKeyCloud    string `toml:"api_key_cloud"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}

func LoadConfig() (*Config, error) {
	config := NewConfig()
	_, err := toml.DecodeFile("./.env", config)
	if err != nil {
		_, err := toml.DecodeFile("./../.env", config)
		if err != nil {
			_, err := toml.DecodeFile("./../../.env", config)
			if err != nil {
				return nil, err
			}
		}
	}

	return config, nil
}
