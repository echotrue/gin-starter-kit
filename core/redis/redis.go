package redis

import (
	"gin-demo/core/config"
	"github.com/gomodule/redigo/redis"
	"log"
	"sync"
)

var (
	instance *RdsConfig
	once     sync.Once
	conn     redis.Conn
	err      error
)

type RdsConfig struct {
	Address string
	Auth    string
}

// Get instance
func Instance() *RdsConfig {
	once.Do(func() {
		c := config.GetConfig().Redis
		instance = &RdsConfig{
			Address: c.Host + ":" + c.Port,
			Auth:    c.Auth,
		}
	})
	return instance
}

// Connect redis
func (rc *RdsConfig) Connect() {
	conn, err = redis.Dial("tcp", rc.Address, redis.DialPassword(rc.Auth))
	if err != nil {
		log.Fatal(err)
	}
}

// Get redis conn
func (rc *RdsConfig) GetRds() redis.Conn {
	return conn
}
