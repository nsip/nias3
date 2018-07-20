package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	N3EngineWebport int
	N3EngineTLSCert string
	Loaded          bool
}

var conf Config

func LoadConfig() Config {
	if !conf.Loaded {
		if _, err := toml.DecodeFile("n3.toml", &conf); err != nil {
			log.Fatalln("Unable to read config file, aborting.", err)
		}
		conf.Loaded = true
	}
	return conf
}
