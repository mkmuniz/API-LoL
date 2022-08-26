package configs

import "github.com/spf13/viper"

var cfg *config

func init() {
	viper.SetDefault("api.port", ":4000")
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", "5432")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
		Host: viper.GetString("api.host"),
	}

	cfg.RiotAPI = APILoL{
		Key: viper.GetString("api.key"),
	}

	return nil
}

func GetRiotAPIKey() string {
	return cfg.RiotAPI.Key
}

func GetAPIConfig() string {
	return cfg.API.Port
}
