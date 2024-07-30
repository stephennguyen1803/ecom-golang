package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		password string `mapstruture:"password"`
		Host     string `mapstructure:"host"`
		DbName   string `mapstructure:"dbName"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config/") // Path to config
	viper.SetConfigName("local")     //File Name
	viper.SetConfigType("yaml")      //Extention

	//read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration: %w", err))
	}

	//read server configuration
	fmt.Println("Server Port::", viper.GetInt("server.port"))
	fmt.Println("Server Port::", viper.GetString("security.jwt.key"))

	//load server configuration using struct
	var config Config
	viper.Unmarshal(&config)
	fmt.Println("Server Structure Port::", config.Server.Port)
	for _, configDatabase := range config.Databases {
		fmt.Println("Database Host ::", configDatabase.Host)
		fmt.Println("Database name ::", configDatabase.DbName)
	}
}
