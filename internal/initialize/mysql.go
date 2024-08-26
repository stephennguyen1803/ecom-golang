package initialize

import (
	"ecom-project/global"
	"ecom-project/internal/po"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	// Mysql initialization
	mysqlConfig := global.Config.Mysql
	// set loc = local to avoid UTC time
	dsn := "%s:%s@tcp(%v:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.DbName)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	if err != nil {
		checkErrorPanic(err, "Failed to connect to database")
	}
	global.Logger.Info("Mysql connected successfully")
	global.Mdb = db

	// set pool
	setPool()
	// create tables
	migrateTables()
	genTableDAO()
}

// open the pool
func setPool() {
	db, err := global.Mdb.DB()
	if err != nil {
		checkErrorPanic(err, "Failed to set pool")
	}
	db.SetMaxIdleConns(global.Config.Mysql.MaxIdleConns)
	db.SetMaxOpenConns(global.Config.Mysql.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(global.Config.Mysql.ConnMaxLifeTime))
}

func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)
	if err != nil {
		checkErrorPanic(err, "Failed to migrate tables")
	}
}

func genTableDAO() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/model",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// // gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(global.Mdb) // reuse your gorm db

	// // Generate basic type-safe DAO API for struct `model.User` following conventions
	// g.ApplyBasic(model.User{})
	g.GenerateModel("go_crm_user")

	// // Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	// g.ApplyInterface(func(Querier) {}, model.User{}, model.Company{})

	// Generate the code
	g.Execute()
}
