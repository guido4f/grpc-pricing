package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Grpc Grpc
}
type Grpc struct {
	Port string `mapstructure:"PORT"`
	Host string `mapstructure:"GRPC_SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("yaml")
	viper.SetConfigName("app")
	viper.AddConfigPath("/app/config/") // path to look for the config file in
	//viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath("test-resources") // optionally look for config in the working directory
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
