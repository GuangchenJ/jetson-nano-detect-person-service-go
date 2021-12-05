package db

import (
	"encoding/json"
	"os"
)

// AppConfig 服务器配置
type AppConfig struct {
	AppName  string   `json:"app_name"`
	Port     string   `json:"port"`
	Mode     string   `json:"mode"`
	DataBase DataBase `json:"data_base"`
	Redis    Redis    `json:"redis"`
	Grpc     Grpc     `json:"grpc"`
}

// DataBase mysql配置
type DataBase struct {
	Drive         string `json:"drive"`
	Port          string `json:"port"`
	User          string `json:"user"`
	Pwd           string `json:"pwd"`
	Host          string `json:"host"`
	Database      string `json:"database"`
	Charset       string `json:"charset"`
	ParseTime     string `json:"parse_time"`
	Loc           string `json:"loc"`
	LogLevel      string `json:"log_level"`
	SlowThreshold int    `json:"slow_threshold"`
}

// Redis redis配置
type Redis struct {
	NetWork string `json:"net_work"`
	Addr    string `json:"addr"`
	Port    string `json:"port"`
	Pwd     string `json:"pwd"`
	Prefix  string `json:"prefix"`
}

type Grpc struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

var ServConfig AppConfig

// InitConfig 初始化服务器配置
func InitConfig() *AppConfig {
	file, err := os.Open("./config.json")
	if nil != err {
		panic(err.Error())
	}

	decoder := json.NewDecoder(file)
	conf := AppConfig{}
	err = decoder.Decode(&conf)
	if nil != err {
		panic(err.Error())
	}

	return &conf
}
