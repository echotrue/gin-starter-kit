package log

import (
	"gin-demo/core/config"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	Logger = logrus.New()
)
var logLevelMap = map[string]logrus.Level{
	"panic": logrus.PanicLevel,
	"fatal": logrus.FatalLevel,
	"error": logrus.ErrorLevel,
	"warn":  logrus.WarnLevel,
	"info":  logrus.InfoLevel,
	"debug": logrus.DebugLevel,
	"trace": logrus.TraceLevel,
}

func New() {
	conf := config.GetConfig()

	Logger.SetLevel(logLevelMap[conf.Log.Level])

	Logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:               false,
		DisableColors:             false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             false,
		TimestampFormat:           "2006-01-02 15:04:05",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		QuoteEmptyFields:          false,
		FieldMap:                  nil,
		CallerPrettyfier:          nil,
	})

	if conf.Log.Target == "File" {
		filePath := conf.Log.File.Filepath + "/" + time.Now().Format("2006-01-02") + ".log"
		Logger.Hooks.Add(NewFileHook(filePath, nil))
	}
}
