package logger

import (
	"fmt"
	"os"
	"strconv"
)

type FileLogger struct {
	level       int
	logPath     string
	logName     string
	file        *os.File
	warnFile    *os.File
	LogDataChan chan *LogData
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
	logChanSize, ok := config["log_chan_size"]
	if !ok {
		logChanSize = "50000"
	}
	chanSize, err := strconv.Atoi(logChanSize)
	if err != nil {
		chanSize = 50000
	}
	level := getLogLevel(logLevel)
	log = &FileLogger{
		level:       level,
		logPath:     logPath,
		logName:     logName,
		LogDataChan: make(chan *LogData, chanSize),
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
	go f.writeLogBackground()
}

func (f *FileLogger) writeLogBackground() {
	// 队列为空时，会阻塞线程，但对主线程没有影响
	for logData := range f.LogDataChan {
		var file *os.File = f.file
		if logData.WarnAndFatal {
			file = f.warnFile
		}
		fmt.Fprintf(file, "[%s] %s (%s:%s:%d) %s\n", logData.TimeStr, logData.LevelStr, logData.FileName, logData.FuncName, logData.LineNo, logData.Message)
	}
}
func (f *FileLogger) DEBUG(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	logData := writeLog(LogLevelDebug, format, args...)
	// 判断队列是否满
	select {
	case f.LogDataChan <- logData:
	default:
	}
}
func (f *FileLogger) TRACE(format string, args ...interface{}) {
	if f.level > LogLevelTrace {
		return
	}
	logData := writeLog(LogLevelTrace, format, args...)
	// 判断队列是否满
	select {
	case f.LogDataChan <- logData:
	default:
	}
}
func (f *FileLogger) INFO(format string, args ...interface{}) {
	if f.level > LogLevelInfo {
		return
	}
	logData := writeLog(LogLevelInfo, format, args...)
	// 判断队列是否满
	select {
	case f.LogDataChan <- logData:
	default:
	}
}
func (f *FileLogger) WARN(format string, args ...interface{}) {
	if f.level > LogLevelWarn {
		return
	}
	logData := writeLog(LogLevelWarn, format, args...)
	// 判断队列是否满
	select {
	case f.LogDataChan <- logData:
	default:
	}
}
func (f *FileLogger) ERROR(format string, args ...interface{}) {
	if f.level > LogLevelError {
		return
	}
	logData := writeLog(LogLevelError, format, args...)
	// 判断队列是否满
	select {
	case f.LogDataChan <- logData:
	default:
	}
}
func (f *FileLogger) FATAL(format string, args ...interface{}) {
	if f.level > LogLevelFatal {
		return
	}
	logData := writeLog(LogLevelFatal, format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
}

func (f *FileLogger) Close() {
	f.file.Close()
	f.warnFile.Close()
}
