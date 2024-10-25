package config

import (
	"encoding/json"
	"os"
)

type wx struct {
	AppiD     string `json:"appid"`
	Appsecret string `json:"appsecret"`
}

type server struct {
	Ip   string `json:"ip"`
	Port uint   `json:"port"`
}

type database struct {
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type config struct {
	Wx       wx       `json:"wx"`
	Server   server   `json:"server"`
	Database database `json:"database"`
	Jwtkey   string   `json:"jwtkey"`
}

var Wx wx
var Server server
var Database database
var Jwtkey string

func init() {
	content, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	var config config
	json.Unmarshal(content, &config)
	Wx = config.Wx
	Server = config.Server
	Database = config.Database
	Jwtkey = config.Jwtkey
}
