package log

import (
	"gin-demo/core/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

// Usage:
// log.Logger.Error("Request error!");
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

	// Request log  and panic recovery
	logPath := conf.Log.File.Filepath
	dir := filepath.Dir(logPath)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		panic("failed create Dir :" + dir + err.Error())
	}
	w, err := os.OpenFile(logPath+"/request.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic("failed create " + logPath + "/request.log," + err.Error())
	}
	gin.DisableConsoleColor()
	gin.DefaultWriter = w
	gin.DefaultErrorWriter = w
}
