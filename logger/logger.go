package logger

import (
	"fmt"
)

var log LogInterface

/*
name:file, console
*/
func InitLogger(name string, config map[string]string) (err error) {
	switch name {
	case "file":
		log, err = NewFileLogger(config)
	case "console":
		log, err = NewConsoleLogger(config)
	default:
		err = fmt.Errorf("unsupport logger name:%s", name)
	}
	return
}

// 封装为包函数，更加容易调用，表现在外面是：能通过包名logger.func来写日志
func DEBUG(format string, args ...interface{}) {
	log.DEBUG(format, args...)
}
func TRACE(format string, args ...interface{}) {
	log.TRACE(format, args...)
}
func INFO(format string, args ...interface{}) {
	log.INFO(format, args...)
}
func WARN(format string, args ...interface{}) {
	log.WARN(format, args...)
}
func ERROR(format string, args ...interface{}) {
	log.ERROR(format, args...)
}
func FATAL(format string, args ...interface{}) {
	log.FATAL(format, args...)
}
