package config

import "os"

// LogFile config
func LogFile() *os.File {
	f, _ := os.OpenFile("log/logfile.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	return f
}
