package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Address string
}

type Config struct {
	//Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HttpServer  `yaml:"http_server"`
}

func MustLoad() *Config {

	var configPath string
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to the configaration file")

		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("config fath is not set")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {

		log.Fatalf("config fath does not exsits, path: %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {

		log.Fatalf("can not read config file, path: %s", err.Error())
	}

	return &cfg
}
