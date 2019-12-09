package core

import (
	"github.com/gin-gonic/gin"
	log "github.com/go-ozzo/ozzo-log"
)

func test() {
	//create the root logger
	logger := log.NewLogger()

	//Set log target
	logFileTarget := log.NewFileTarget()
	logFileTarget.FileName = ""
	logFileTarget.MaxLevel = log.LevelEmergency
	logger.Targets = append(logger.Targets, logFileTarget)

	logger.Open()

	defer logger.Close()

}

type GinLogger struct {
	engine  *gin.Engine
	FilePath string
}

// NewGinLogger create a gin root logger
func NewGinLogger(r *gin.Engine) *GinLogger {
	return &GinLogger{
		engine: r,
	}
}

func (gl *GinLogger) NewFileTarget() {

}

type Target struct {
}
