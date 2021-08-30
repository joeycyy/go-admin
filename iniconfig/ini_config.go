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

	// fmt.Println(lineArr)
	// 获取result的反射类型对象
	ptrTypeObj := reflect.TypeOf(result)
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
		// fmt.Println(line)
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
			lastSectionName, err = parseSectionName(line, elemTypeObj)
			// fmt.Println(lastSectionName)
			if err != nil {
				err = fmt.Errorf("%v, lineno:%d", err, index+1)
				return
			}
			continue
		}
		// 如果是节点字段行
		// err = parseItemName(line, lastSectionName, result)

	}
	_ = lastSectionName
	return
}

// 解析字段名称
func parseItemName(line string, sectionName string, result interface{}) (err error) {
	index := strings.Index(line, "=")
	// 判断是否有等号
	if index == -1 {
		err = fmt.Errorf("syntax error, line:%s", line)
		return
	}

	// 获取左边的key值
	key := strings.TrimSpace(line[0:index])
	// 获取右边的value值
	value := strings.TrimSpace(line[index+1:])

	_ = value
	// 判断key值的长度
	if len(key) == 0 {
		err = fmt.Errorf("syntax error, line:%s", line)
		return
	}

	// 获取节点的类型对象
	valueObj := reflect.ValueOf(result)
	sectionValue := valueObj.Elem().FieldByName(sectionName)
	sectionType := sectionValue.Type()

	if sectionType.Kind() != reflect.Struct {
		err = fmt.Errorf("%s must be struct", sectionName)
		return
	}

	return
}

// 解析节点名称，返回节点名称
func parseSectionName(line string, TypeObj reflect.Type) (fieldName string, err error) {
	// 判断节点格式是否正确
	if line[len(line)-1] != ']' {
		err = fmt.Errorf("syntax error, invalid section:%s", line)
		return
	}
	if line[0] == '[' && line[len(line)-1] == ']' {
		// 获取中括号内的内容
		sectionName := strings.TrimSpace(line[1 : len(line)-1])
		// 判断中括号内的内容是否为空
		if len(sectionName) == 0 {
			err = fmt.Errorf("syntax error, invalid section:%s", line)
			return
		}
		// 遍历类型对象，获取其字段名称
		for i := 0; i < TypeObj.NumField(); i++ {
			field := TypeObj.Field(i)
			tagValue := field.Tag.Get("ini")
			if tagValue == sectionName {
				fieldName = field.Name
				break
			}
		}
	}
	return
}
