package global

import (
	logger "ecom-project/pkg/logger"
	"ecom-project/pkg/settings"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
)
