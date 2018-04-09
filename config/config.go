package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

)


type LogConfig struct {
	Access string `json:"access"`
	Error  string `json:"error"`
}

type AuditorConfig struct {
	Sms  []string `json:"sms"`
	Mail []string `json:"mail"`
}

type HttpConfig struct {
	Listen string `json:"listen"`
	Secret string `json:"secret"`
}

type RpcConfig struct {
	Listen string `json:"listen"`
}

type MysqlConfig struct {
	Addr    string 			  `json:"addr"`
	Idle    int               `json:"idle"`
	Max     int               `json:"max"`
	ShowSQL bool              `json:"sql"`
}

type TimeoutConfig struct {
	Conn  int64 `json:"conn"`
	Read  int64 `json:"read"`
	Write int64 `json:"write"`
}

type RedisConfig struct {
	Addr    string         `json:"addr"`
	PicOvertime int64 		`json:"overtime"`   // 验证码超时时间
	Idle    int            `json:"idle"`
	Max     int            `json:"max"`
	Timeout *TimeoutConfig `json:"timeout"`
}



type GlobalConfig struct {
	Debug   	bool                `json:"debug"`   //
	Token   	string              `json:"token"` 	 //
	Request     string 				`json:"request"` // 是否使用https
	Salt    	string              `json:"salt"`    // 验证key
	Log     	*LogConfig          `json:"log"`
	Http    	*HttpConfig         `json:"http"`
	Mysql   	*MysqlConfig        `json:"mysql"`
	Redis   	*RedisConfig        `json:"redis"`
}

var (
	File string
	G    *GlobalConfig
)

func Parse(cfg string) error {
	if cfg == "" {
		return fmt.Errorf("使用 -c 添加导入配置文件")
	}

	if !fileExists(cfg) {
		return fmt.Errorf("配置文件 %s 不存在", cfg)
	}

	File = cfg

	configContent, err := ioutil.ReadFile(cfg)
	if err != nil {
		return fmt.Errorf("读取配置文件 %s 错误 %s", cfg, err.Error())
	}

	var c GlobalConfig
	err = json.Unmarshal(configContent, &c)
	if err != nil {
		return fmt.Errorf("解析配置文件 %s 错误 %s", cfg, err.Error())
	}

	G = &c

	log.Println("加载配置文件", cfg, "成功")
	return nil
}

func fileExists(fp string) bool {
	_, err := os.Stat(fp)
	return err == nil || os.IsExist(err)
}
