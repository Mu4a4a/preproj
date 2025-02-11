package config

import (
	"github.com/spf13/viper"
	"time"
)

type RedisConfig struct {
	Addr        string        `yaml:"addr"`
	Password    string        `yaml:"password"`
	User        string        `yaml:"g"`
	DB          int           `yaml:"db"`
	MaxRetries  int           `yaml:"max_retries"`
	DialTimeout time.Duration `yaml:"dial_timeout"`
	Timeout     time.Duration `yaml:"timeout"`
}

func LoadRedisConfig() RedisConfig {
	return RedisConfig{
		Addr:        viper.GetString("cache.addr"),
		Password:    viper.GetString("cache.password"),
		User:        viper.GetString("cache.user"),
		DB:          viper.GetInt("cache.db"),
		MaxRetries:  viper.GetInt("cache.maxRetries"),
		DialTimeout: viper.GetDuration("cache.dialTimeout"),
		Timeout:     viper.GetDuration("cache.timeout"),
	}
}
