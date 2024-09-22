package global

import (
	"database/sql"
	"ecom-project/pkg/logger"
	"ecom-project/pkg/settings"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Mdb    *gorm.DB
	Mdbc   *sql.DB
	Config settings.Config
	Logger *logger.LoggerZap
	Redis  *redis.Client
)
