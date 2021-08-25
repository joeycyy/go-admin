package logger

import (
	"fmt"
	"os"
)

type ConsoleLogger struct {
	level int
}

func NewConsoleLogger(config map[string]string) (log LogInterface, err error) {
	logLevel, ok := config["log_level"]
	if !ok {
		err = fmt.Errorf("not found log_level")
		return
	}
	level := getLogLevel(logLevel)
	log = &ConsoleLogger{
		level: level,
	}
	return
}
func (c *ConsoleLogger) Init() {

}

func (c *ConsoleLogger) DEBUG(format string, args ...interface{}) {
	if c.level > LogLevelDebug {
		return
	}
	writeLog(os.Stdout, LogLevelDebug, format, args...)
}
func (c *ConsoleLogger) TRACE(format string, args ...interface{}) {
	if c.level > LogLevelTrace {
		return
	}
	writeLog(os.Stdout, LogLevelDebug, format, args...)
}
func (c *ConsoleLogger) INFO(format string, args ...interface{}) {
	if c.level > LogLevelInfo {
		return
	}
	writeLog(os.Stdout, LogLevelInfo, format, args...)
}
func (c *ConsoleLogger) WARN(format string, args ...interface{}) {
	if c.level > LogLevelWarn {
		return
	}
	writeLog(os.Stdout, LogLevelDebug, format, args...)
}
func (c *ConsoleLogger) ERROR(format string, args ...interface{}) {
	if c.level > LogLevelError {
		return
	}
	writeLog(os.Stdout, LogLevelDebug, format, args...)
}
func (c *ConsoleLogger) FATAL(format string, args ...interface{}) {
	if c.level > LogLevelFatal {
		return
	}
	writeLog(os.Stdout, LogLevelFatal, format, args...)
}

func (c *ConsoleLogger) Close() {
}
