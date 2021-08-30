package iniconfig

import (
	"fmt"
	"reflect"
	"strings"
)

// data 字节数组，可以是读取文件时获取的字节数组
// result 空接口，反射对象的种类要为指针
// 该函数可以将字节数组转换为对应的struct
func UnMarshal(data []byte, result interface{}) (err error) {
	// 获取文本行数组
	lineArr := strings.Split(string(data), "\n")

	fmt.Println(lineArr)
	// 获取result的反射类型对象
	ptrTypeObj := reflect.TypeOf(result)
	fmt.Println(ptrTypeObj)
	// 判断种类是否为指针
	if ptrTypeObj.Kind() != reflect.Ptr {
		err = fmt.Errorf("please pass address")
		return
	}
	// 判断指针指向的元素类型是否为struct
	elemTypeObj := ptrTypeObj.Elem()
	if elemTypeObj.Kind() != reflect.Struct {
		err = fmt.Errorf("please pass struct")
		return
	}

	var lastSectionName string
	// 开始按行处理文本内容
	for index, line := range lineArr {
		line = strings.TrimSpace(line)
		// 忽略空行
		if len(line) == 0 {
			continue
		}
		// 忽略注释行
		if line[0] == ';' || line[0] == '#' {
			continue
		}
		// 判断开头首字符是否为'['
		// 如果是节点名称行
		if line[0] == '[' {
			lastSectionName = parseSectionName(line)
			continue
		}
		// 如果是节点字段行
		parseItemName(line, lastSectionName, result)

	}
	return
}

// 解析节点名称行，返回节点名称
func parseSectionName(line string, result interface{}) (lastSectionName string, err error) {
	// 判断节点格式是否正确
	if line[len(line)-1] != ']' {
		err = fmt.Errorf("syntax error, line:%s", line)
	}
	return
}
