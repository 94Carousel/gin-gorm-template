package config

import (
	"os"

	"github.com/gin-gonic/gin"
)

// LogFile config
func LogFile() *os.File {
	logName := gin.Mode() + ".log"
	f, _ := os.OpenFile("log/"+logName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	return f
}
