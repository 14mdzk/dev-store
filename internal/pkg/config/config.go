package config

import "github.com/spf13/viper"

type Config struct {
	DBDriver     string `mapstruct:"DB_DRIVER"`
	DBConnection string `mapstruct:"DB_CONNECTION"`
	ServerPort   string `mapstruct:"SERVER_PORT"`
}

func LoadConfig(fileConfigPath string) (Config, error) {
	var config Config

	viper.AddConfigPath(fileConfigPath)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	viper.Unmarshal(config)
	return config, nil
}
