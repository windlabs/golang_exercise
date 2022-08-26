package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// ini 文件解析器
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

//Config
type Config struct {
	MysqlConfig `ini:"mysql"`
}

func loadIni(fileName string, data interface{}) (err error) {
	//1.0参数的校验
	//1.1传进来的data参数必须是指针类型(因为需要在函数中对其赋值，值类型/引用类型的区别)
	t := reflect.TypeOf(data)
	fmt.Printf("测试:%v  %v\n", t, t.Kind())
	if t.Kind() != reflect.Ptr {
		err = errors.New("data should be a pointer.")
		return
	}

	//1.2传进来的data参数必须是结构体类型(配置文件不止一个字段)
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data should be a Struct pointer.")
		return
	}

	//2.0读文件得到字节类型数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}

	//string(b) //将字节数组转换成字符串
	lineSlice := strings.Split(string(b), "\r\n")
	//fmt.Printf("%#v\n", lineSlice)
	structName := ""

	//3.0一行一行读数据
	for idx, line := range lineSlice {
		//去掉字符串首尾空格
		line = strings.TrimSpace(line)
		//如果是空行 直接过滤
		if len(line) == 0 {
			continue
		}

		//3.1如果是注释就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}

		//3.2如果是[]就表示是节(section)
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			//[]  [...]若...为空也不行
			//[]中间内容
			strings.TrimSpace(line[1 : len(line)-1])
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			//根据字符串sectionName的内容去结构体中找data里面根据反射找对应的结构体
		} else {
			//3.3如果不是[]开头，就是内容，  host=127.0.0.1  用=分割键值对
			//按照等号分割这一行 等号左边是key  等号右边是value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}

			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])

			//根据structName去data里面把对应的嵌套结构体给取出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) //拿到嵌套结构体的值信息
			sType := sValue.Type()                     //拿到嵌套结构体的类型信息
			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data中的%s字段应该是一个结构体", structName)
				return
			}

			var fieldName string
			var fileType reflect.StructField
			//遍历嵌套结构体的每一个字段 判断tag是不是等于key
			for i := 0; i < sValue.NumField(); i++ {
				filed := sType.Field(i)
				//tag信息是存储在类型信息中的
				fileType = filed
				if filed.Tag.Get("ini") == key {
					//找到对应的字段
					fieldName = filed.Name
					break
				}
			}
			//如果key=tag 给这个字段赋值
			//根据fieldName去取出这个字段
			if len(fieldName) == 0 {
				//在结构体中找不到对应的字符
				continue
			}

			fileObj := sValue.FieldByName(fieldName)
			//对其赋值
			fmt.Println(fieldName, fileType.Type.Kind())
			switch fileType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetBool(valueBool)
			}
		}
	}
	return
}

func main() {
	mc := Config{}
	err := loadIni("./config.ini", &mc)
	if err != nil {
		fmt.Printf("load error %v\n", err)
	}
	fmt.Println(mc.Address, mc.Port, mc.Username, mc.Password)

}
