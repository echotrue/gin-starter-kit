package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

const (
	LevelEmergency Level = iota
	LevelAlert
	LevelCritical
	LevelError
	LevelWarning
	LevelNotice
	LevelInfo
	LevelDebug
)

type Level int

var LevelNames = map[Level]string{
	LevelDebug:     "DEBUG",
	LevelInfo:      "INFO",
	LevelNotice:    "NOTICE",
	LevelWarning:   "WARNING",
	LevelError:     "ERROR",
	LevelCritical:  "CRITICAL",
	LevelAlert:     "ALERT",
	LevelEmergency: "EMERGENCY",
}

var (
	err    error
	Logger *LocalLog
)

type LocalLog struct {
	fileTarget *os.File
	engine     *gin.Engine
}

// New log
func New(r *gin.Engine) *LocalLog {
	return &LocalLog{
		fileTarget: nil,
		engine:     r,
	}
}

// init log
func (l *LocalLog) FileTarget() {
	// 创建日志
	l.createFile()
	//日志配置
	c := l.logConfig()
	// 设置默认日志引擎
	gin.DefaultWriter = io.MultiWriter(l.fileTarget)
	//使用自定义日志中间件
	l.engine.Use(gin.LoggerWithConfig(c))
	//异常信息写入日志
	l.engine.Use(gin.Recovery())

	Logger = l
}

// Create file
func (l *LocalLog) createFile() {
	filename := time.Now().Format("20060102")
	l.fileTarget, err = os.OpenFile("./runtime/log/app."+filename+".log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic("Create log failed!")
	}
}

func (l *LocalLog) logConfig() gin.LoggerConfig {
	c := gin.LoggerConfig{}
	c.Formatter = func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%v [%v][%v] %v %v %v \n",
			params.TimeStamp.Format("2006-01-02 15:04:05"), params.Method, LevelNames[LevelDebug], params.Path, params.ClientIP, params.ErrorMessage)
		//return fmt.Sprintf("%s \n", params.Method)
	}
	c.Output = io.MultiWriter(l.fileTarget)
	return c
}

func (l *LocalLog) Log(level Level, m ...interface{}) {
	msg := ""
	for _, v := range m {
		msg += fmt.Sprintf("%v", v)
	}
	message := fmt.Sprintf("%v [%v]: %v \n", time.Now().Format("2006-01-02 15:04:05"), LevelNames[level], msg)
	_, err = fmt.Fprint(l.fileTarget, message)
}
func (l *LocalLog) Error(m ...interface{}) {
	l.Log(LevelError, m...)
}
func (l *LocalLog) Warning(m ...interface{}) {
	l.Log(LevelWarning, m...)
}
func (l *LocalLog) Debug(m ...interface{}) {
	l.Log(LevelDebug, m...)
}
func (l *LocalLog) Info(m ...interface{}) {
	l.Log(LevelInfo, m...)
}
func (l *LocalLog) Notice(m ...interface{}) {
	l.Log(LevelNotice, m...)
}
