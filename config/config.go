package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

var config Config

type Config struct {
	Server *Server `json:"server" yaml:"server"`
	DB     *DB     `json:"db" yaml:"db"`
}

type Server struct {
	Port string `json:"port" yaml:"port"`
}

type DB struct {
	Dialect    string `json:"dialect" yaml:"dialect"`
	DataSource string `json:"datasource" yaml:"datasource"`
}

func Init(path string) {
	rawConfig, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("unable to read config file")
		return
	}

	err = yaml.Unmarshal(rawConfig, &config)
	if err != nil {
		log.Fatal("unable to unmarshal config file")
		return
	}
}

func Get() Config {
	return config
}
