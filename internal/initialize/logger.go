package initialize

import (
	"ecom-project/global"
	"ecom-project/pkg/logger"
	"fmt"
)

func InitLogger() {
	// Init Logger
	fmt.Printf("Logger File Name: %v\n", global.Config.Logger.FileLogName)
	global.Logger = logger.NewLogger(global.Config.Logger)
}
