package config

import (
	/* "encoding/json"
	"fmt"
	"io/ioutil"
	"os" */
)
/* 
func init() {
	fmt.Print("Load Config ") //在主程序运行之初读取配置文件
	// 如果config中没有对应的值，则会在setting.go中寻找
	ConfigData = func(src string) *DATA {
		data, _ := ioutil.ReadFile(src)
		newDATA := &DATA{}
		json.Unmarshal(data, &newDATA)
		if ConfigRemove {
			os.Remove(src)
		}
		return newDATA
	}(ConfigPath)
}

// 获取名为name的JWT密钥
func JWTKEY(name string) string {
	return JWTMAP[name].Key
}

// 获取名为name的JWT标签
func JWTTAG(name string) string {
	return JWTMAP[name].Tag
} */

