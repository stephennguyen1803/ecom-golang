package initialize

import (
	"ecom-project/global"
	"fmt"
	"strconv"

	"go.uber.org/zap"
)

func Run() {
	//1 - load config
	LoadConfig()
	//2 - init logger
	InitLogger()
	global.Logger.Info("Logger init success 1234", zap.String("logger", "zap"))
	//3 - init mysql
	InitMysql()
	fmt.Printf("mysql.username: %v\n", global.Config.Mysql.User)
	//4 - init redis
	InitRedis()
	//5 - init router
	r := InitRouter()
	port := ":" + strconv.Itoa(global.Config.Server.Port)
	r.Run(port)
}
