package global

import (
	"ecom-project/pkg/logger"
	"ecom-project/pkg/settings"

	"gorm.io/gorm"
)

var (
	Mdb    *gorm.DB
	Config settings.Config
	Logger *logger.LoggerZap
)
