package iniconfig

import (
	"fmt"
	"testing"
)

type Config struct {
	MysqlConf  MysqlConfig  `ini:"mysql"`
	ServerConf ServerConfig `ini:"server"`
}
type MysqlConfig struct {
	UserName string `ini:"username"`
	Passwd   string `ini:"passwd"`
	Database string `ini:"database"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
}

type ServerConfig struct {
	Ip   string `ini:"ip"`
	Port int    `ini:"port"`
}

func TestIniConfig(t *testing.T) {
	var conf = &Config{}
	err := UnMarshalFile("./config.ini", conf)
	if err != nil {
		t.Errorf("UnMarshal failed, err:%s", err)
		return
	}
	fmt.Printf("conf:%#v", *conf)
}
