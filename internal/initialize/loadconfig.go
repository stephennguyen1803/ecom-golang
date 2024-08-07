package initialize

import (
	"ecom-project/global"
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() {
	// Load configuration from file
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
	if err = viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("failed to unmarshal configuration: %w", err))
	}
	// fmt.Println("Server Structure Port::", config.Port)
	// for _, configDatabase := range global.Config.Databases {
	// 	fmt.Println("Database Host ::", configDatabase.Host)
	// 	fmt.Println("Database name ::", configDatabase.DbName)
	// }
}
