package util

import "github.com/sirupsen/logrus"

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	Logger.SetFormatter(&logrus.TextFormatter{
		// 时间戳的格式
		TimestampFormat: "2006-01-02 15:04:05",
	})
	Logger.SetLevel(logrus.DebugLevel)
}
