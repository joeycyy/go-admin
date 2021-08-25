package logger

import (
	"testing"
)

func TestFileLogger(t *testing.T) {
	log, _ := NewFileLogger(map[string]string{
		"log_path":  "c:/log/",
		"log_name":  "test",
		"log_level": "debug",
	})
	log.DEBUG("userid[%d] is from china", 911)
	log.INFO("userid[%d] is from china", 911)

	log2, _ := NewConsoleLogger(map[string]string{
		"log_path":  "c:/log/",
		"log_name":  "test",
		"log_level": "debug",
	})
	log2.DEBUG("userid[%d] is from china", 911)
	log2.INFO("userid[%d] is from china", 9101)
}
