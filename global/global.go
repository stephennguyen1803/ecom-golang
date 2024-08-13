package global

import (
	"ecom-project/pkg/logger"
	"ecom-project/pkg/settings"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Mdb    *gorm.DB
	Config settings.Config
	Logger *logger.LoggerZap
	Redis  *redis.Client
)
