package config

import (
	"os"

	"github.com/naoina/toml"
)

// 환경구성
// env -> toml

type Config struct {
	Server struct {
		Port string
	}
}

// config load function
func NewConfig(filtPath string) *Config {
	c := new(Config)

	if file, err := os.Open(filtPath); err != nil {
		panic(err)
	} else if err = toml.NewDecoder(file).Decode(c); err != nil {
		panic(err)
	} else {
		return c
	}
}
