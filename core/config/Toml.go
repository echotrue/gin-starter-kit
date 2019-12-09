package config

import (
	"errors"
	"github.com/BurntSushi/toml"
	"log"
	"sync"
)

var (
	ParseErr = errors.New("解析配置文件失败")
)

var (
	once       sync.Once
	instance   *Toml
	configData *Config
)

// Toml struct
type Toml struct {
	path string
}

// Create config Object
func NewToml() *Toml {
	once.Do(func() {
		instance = &Toml{
			path: "./application.toml",
		}
	})
	return instance
}

// Parse Toml file
func (t *Toml) Parse() {
	_, err := toml.DecodeFile(t.path, &configData)
	if err != nil {
		log.Fatal(ParseErr, err)
	}
}

// Get config data
// Return config struct
func GetConfig() *Config {
	return configData
}
