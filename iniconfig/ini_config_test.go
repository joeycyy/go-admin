package iniconfig

import (
	"io/ioutil"
	"testing"
)

type Config struct {
}
type MysqlConfig struct {
}

type ServerConfig struct {
}

func TestIniConfig(t *testing.T) {
	data, err := ioutil.ReadFile("./config.ini")
	if err != nil {
		t.Errorf("read file failed, err:%s", err)
	}
	err = UnMarshal(data, &Config{})
	if err != nil {
		t.Errorf("UnMarshal failed, err:%s", err)
		return
	}
}
