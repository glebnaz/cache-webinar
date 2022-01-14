package config

import "github.com/kelseyhightower/envconfig"

//BaseConfig is configs with base param
type BaseConfig struct {
	HTTPPort  string `json:"http_port" default:"8080"`
	RedisHost string `json:"redis_host"`
}

type WriterConfig struct {
	BaseConfig
}

type ReaderConfig struct {
	BaseConfig
}

func NewReaderConfig() (ReaderConfig, error) {
	var cfg ReaderConfig
	err := envconfig.Process("READER", &cfg)
	return cfg, err
}

func NewWriteConfig() (WriterConfig, error) {
	var cfg WriterConfig
	err := envconfig.Process("WRITER", &cfg)
	return cfg, err
}
