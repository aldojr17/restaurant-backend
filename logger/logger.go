package logger

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/logger"
)

func createLogger() logger.Interface {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, //Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	return newLogger
}

func NewGormLogger() logger.Interface {
	return createLogger()
}

func GetGinCustomLogger(params gin.LogFormatterParams) string {
	return fmt.Sprintf("[Info] [%s] [Code: %d] \"%s %s\" %s %s\n\n",
		params.TimeStamp.Format("15:04:05"),
		params.StatusCode,
		params.Method,
		params.Path,
		params.Latency,
		params.ErrorMessage,
	)
}
