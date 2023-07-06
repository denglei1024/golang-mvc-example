package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName string `json:"appname"`
	Port    string `json:"port"`
	Host    string `json:"host"`
}

var _cfg *Config = nil

func ParseConfig(name string) (*Config, error) {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err = decoder.Decode(&_cfg); err != nil {
		return nil, err
	}
	return _cfg, nil
}
