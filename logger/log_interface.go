package logger

type LogInterface interface {
	Init()
	DEBUG(format string, args ...interface{})
	TRACE(format string, args ...interface{})
	INFO(format string, args ...interface{})
	WARN(format string, args ...interface{})
	ERROR(format string, args ...interface{})
	FATAL(format string, args ...interface{})
	Close()
}
