package main

import (
	"go-admin/logger"
	"time"
)

func initLogger(name, logPath, logName string, level string) (err error) {
	config := make(map[string]string, 8)
	config["log_path"] = logPath
	config["log_name"] = logName
	config["log_level"] = level
	err = logger.InitLogger(name, config)
	if err != nil {
		return
	}

	logger.DEBUG("init logger success")
	return
}

func Run() {
	for {
		logger.DEBUG("user server is running")
		logger.INFO("info log record")
		time.Sleep(time.Second)
	}
}

func main() {
	initLogger("file", "c:/log/", "user_server", "debug") // 后续放到配置文件中
	Run()
	return
	//TODO 引入路由，启动http服务
}
