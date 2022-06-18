package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"

	"github.com/pkg/errors"
)

type Config interface {
	Logger
	Listener
	Databaser
}

type config struct {
	Logger
	Listener
	Databaser
}

type envConfig struct {
	LogLevel string            `yaml:"log_level"`
	Address  string            `yaml:"address"`
	Database envDatabaseConfig `yaml:"database"`
}

func New(path string) Config {
	cfg := envConfig{}

	yamlConfig, err := ioutil.ReadFile(path)
	if err != nil {
		panic(errors.Wrapf(err, "failed to read config %s", path))
	}

	err = yaml.Unmarshal(yamlConfig, &cfg)
	if err != nil {
		panic(errors.Wrapf(err, "failed to unmarshal config %s", path))
	}

	return &config{
		Logger:    NewLogger(cfg.LogLevel),
		Listener:  NewListener(cfg.Address),
		Databaser: NewDatabaser(cfg.Database.toPSQLPath(), cfg.Database.Driver),
	}
}
