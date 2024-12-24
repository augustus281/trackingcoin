package global

import (
	database "github.com/augustus281/trackingcoin/database/sqlc"
	"github.com/augustus281/trackingcoin/pkg/logger"
	"github.com/augustus281/trackingcoin/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Db     *database.Store
)
