package configs

import "github.com/sirupsen/logrus"

const (
	PORT         = 8081
	LOGS_DIR     = "./logs/"
	LOGFILE_NAME = "server.log"
	LOG_LEVEL    = logrus.DebugLevel
)
