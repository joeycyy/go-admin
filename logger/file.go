package logger

import (
	"fmt"
	"os"
)

type FileLogger struct {
	level    int
	logPath  string
	logName  string
	file     *os.File
	warnFile *os.File
}

func NewFileLogger(config map[string]string) (log LogInterface, err error) {
	logPath, ok := config["log_path"]
	if !ok {
		err = fmt.Errorf("not found log_path")
		return
	}
	logName, ok := config["log_name"]
	if !ok {
		err = fmt.Errorf("not found log_name")
		return
	}
	logLevel, ok := config["log_level"]
	if !ok {
		err = fmt.Errorf("not found log_level")
		return
	}
	level := getLogLevel(logLevel)
	log = &FileLogger{
		level:   level,
		logPath: logPath,
		logName: logName,
	}
	log.Init()
	return
}
func (f *FileLogger) Init() {
	file, err := os.OpenFile(fmt.Sprintf("%s%s.log", f.logPath, f.logName), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)

	if err != nil {
		panic(fmt.Sprintf("open file failed, err:%v", err))
	}

	f.file = file

	file, err = os.OpenFile(fmt.Sprintf("%s%sinfo.log", f.logPath, f.logName), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)

	if err != nil {
		panic(fmt.Sprintf("open file failed, err:%v", err))
	}

	f.warnFile = file
}

func (f *FileLogger) DEBUG(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	writeLog(f.file, LogLevelDebug, format, args...)
}
func (f *FileLogger) TRACE(format string, args ...interface{}) {
	if f.level > LogLevelTrace {
		return
	}
	writeLog(f.file, LogLevelTrace, format, args...)
}
func (f *FileLogger) INFO(format string, args ...interface{}) {
	if f.level > LogLevelInfo {
		return
	}
	writeLog(f.warnFile, LogLevelDebug, format, args...)
}
func (f *FileLogger) WARN(format string, args ...interface{}) {
	if f.level > LogLevelWarn {
		return
	}
	writeLog(f.warnFile, LogLevelWarn, format, args...)
}
func (f *FileLogger) ERROR(format string, args ...interface{}) {
	if f.level > LogLevelError {
		return
	}
	writeLog(f.warnFile, LogLevelError, format, args...)
}
func (f *FileLogger) FATAL(format string, args ...interface{}) {
	if f.level > LogLevelFatal {
		return
	}
	writeLog(f.warnFile, LogLevelFatal, format, args...)
}

func (f *FileLogger) Close() {
	f.file.Close()
	f.warnFile.Close()
}
