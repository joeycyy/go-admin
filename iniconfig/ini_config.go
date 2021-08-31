package iniconfig

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

func UnMarshalFile(filename string, result interface{}) (err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	return UnMarshal(data, result)
}

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
		err = parseItemName(line, lastSectionName, result)
		if err != nil {
			err = fmt.Errorf("%v, lineno:%d", err, index+1)
			return
		}

	}
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

	// 获取result的类型对象
	resultValueObj := reflect.ValueOf(result)
	// 获取节点的value对象
	sectionValue := resultValueObj.Elem().FieldByName(sectionName)
	// 获取节点对应的类型对象
	sectionType := sectionValue.Type()

	// fmt.Println(sectionType)
	// 判断是否结构体类型
	if sectionType.Kind() != reflect.Struct {
		err = fmt.Errorf("%s must be struct", sectionName)
		return
	}

	keyFieldName := ""
	// 遍历MysqlConfig所拥有的字段，判断其tag是否和当前文本行相等,相等则获取其变量名
	for i := 0; i < sectionType.NumField(); i++ {
		field := sectionType.Field(i)
		tagVal := field.Tag.Get("ini")
		if key == tagVal {
			keyFieldName = field.Name
			break
		}
	}

	// 判断keyFieldName是否存在
	if len(keyFieldName) == 0 {
		return
	}
	// 获取具体字段的反射对象
	fieldValue := sectionValue.FieldByName(keyFieldName)
	// fmt.Println(fieldValue)
	if fieldValue == reflect.ValueOf(nil) {
		err = fmt.Errorf("syntax error,line:%s", line)
		return
	}

	// 根据字段类型，设置值
	switch fieldValue.Type().Kind() {
	case reflect.String:
		fieldValue.SetString(value)
	case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int32, reflect.Int64:
		intVal, errRet := strconv.ParseInt(value, 10, 64)
		if errRet != nil {
			err = errRet
			return
		}
		fieldValue.SetInt(intVal)
	case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint32, reflect.Uint64:
		intVal, errRet := strconv.ParseUint(value, 10, 64)
		if errRet != nil {
			err = errRet
			return
		}
		fieldValue.SetUint(intVal)
	case reflect.Float32, reflect.Float64:
		floatVal, errRet := strconv.ParseFloat(value, 64)
		if errRet != nil {
			return
		}

		fieldValue.SetFloat(floatVal)

	default:
		err = fmt.Errorf("unsupport type:%v", fieldValue.Type().Kind())
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
