package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	configPath = ".env"
)

type Config struct {
	Pipeline        PipelineConfig
	MongoDB         MongoConfig
	ClearCollection bool
}

type PipelineConfig struct {
	SourceDir    string
	ChanLen      int
	StageWorkers int
}

type MongoConfig struct {
	User     string `env:"MONGO_INITDB_ROOT_USERNAME"`
	Password string `env:"MONGO_INITDB_ROOT_PASSWORD"`
	DB       string `env:"MONGO_INITDB_DATABASE"`
	Host     string `env:"MONGO_INITDB_HOST"`
	Port     string `env:"MONGO_INITDB_PORT"`
}

func MustLoad() *Config {
	config := &Config{}

	err := cleanenv.ReadConfig(configPath, config)
	if err != nil {
		log.Fatalf("Error while loading config: %s", err)
	}

	return config
}
