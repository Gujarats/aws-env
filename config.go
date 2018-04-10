package main

import (
	"fmt"
	"log"
	"os/user"
	"path"

	"github.com/spf13/viper"
)

type Config struct {
	Profile       string `viper:"profile"`
	AwsConfigPath string `viper:awsConfigPath`
}

const (
	pathConfig = ".aws-env"
)

func getHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

func getConfig() *Config {
	viper.AddConfigPath("$HOME/" + pathConfig)
	viper.SetConfigName("config")
	viper.SetDefault("profile", "default")
	viper.SetDefault("AwsConfigPath", path.Join(getHomeDir(), ".aws/credentials"))
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file must exist in ~/"+pathConfig+"./config.yaml: %s \n", err))
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		panic(fmt.Errorf("Fatal error unmarhsal config struct : %s \n", err))
	}

	return config
}
