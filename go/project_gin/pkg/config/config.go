package config

import (
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v2"
)

// CacheOptions 读取配置文件中cache段，并写入CacheOptions
type CacheOptions struct {
	Backend string        `yaml:"backend"`
	Expire  time.Duration `yaml:"expire"`
	Cleanup time.Duration `yaml:"cleanup"`
}

// HTTPOptions 配置Http服务
type HTTPOptions struct {
	Host   string `yaml:"host"`
	Prefix string `yaml:"prefix"`
	Debug  bool   `yaml:"debug"`
}

type LogOptions struct {
	Format       string `yaml:"format" default:"text"`
	DisableColor bool   `yaml:"disable_color" default:"true"`
	Level        string `yaml:"level" default:"info"`
	Filename     string `yaml:"filename" default:"./app.log"`
}

// Config App配置文件信息
type Config struct {
	HTTPConfig  *HTTPOptions  `yaml:"http"`
	CacheConfig *CacheOptions `yaml:"cache"`
	LogConfig   *LogOptions   `yaml:"log"`
}

var (
	GlobalConfig *Config
)

// New 通过config path新建一个Config对象
func New(configFilePath *string) (config *Config, err error) {
	GlobalConfig = &Config{}
	yamlFile, err := ioutil.ReadFile(*configFilePath)
	if err != nil {
		log.Printf("configFile Read fail #%v ", err.Error())
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, GlobalConfig)
	if err != nil {
		log.Fatalf("configFile Unmarshal fail: %v", err.Error())
		return nil, err
	}
	return GlobalConfig, nil
}
