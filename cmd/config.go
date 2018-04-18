package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/Gujarats/logger"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

type Config struct {
	Profile       string `viper:"profile"`
	AwsConfigPath string `viper:awsConfigPath`
}

const (
	pathConfig = ".aws-env"
)

func getConfig() *Config {
	viper.AddConfigPath("$HOME/" + pathConfig)
	viper.SetConfigName("config")

	home, err := homedir.Dir()
	if err != nil {
		logger.Debug("Error", err)
		os.Exit(1)
	}

	// default value config
	viper.SetDefault("AwsConfigPath", path.Join(home, ".aws/credentials"))
	viper.SetDefault("profile", "default")
	viper.AutomaticEnv() // read in environment variables that match

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file must exist in ~/"+pathConfig+"./config.yaml: %s \n", err))
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		panic(fmt.Errorf("Fatal error unmarhsal config struct : %s \n", err))
	}

	return config
}
