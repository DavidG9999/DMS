package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env           string        `yaml:"env" env-default:"local"`
	TokenTTL      time.Duration `yaml:"token_ttl" env-default:"1h"`
	GRPC          GRPCConfig    `yaml:"grpc"`
	UserClient    UserClient    `yaml:"user_client"`
	}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type UserClient struct {
	Address      string        `yaml:"address"`
	Timeout      time.Duration `yaml:"timeout"`
	RetriesCount int           `yaml:"retries_count"`
}


func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}

	return MustLoadPath(configPath)
}

func MustLoadPath(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file dose not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}
	return &cfg
}

func fetchConfigPath() string {
	var res string
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		if err := godotenv.Load(); err != nil {
			panic("failed to load env var: " + err.Error())
		}
		res = os.Getenv("CONFIG_PATH")
	}

	if res == "test" {
		if err := godotenv.Load(); err != nil {
			panic("failed to load env var: " + err.Error())
		}
		res = os.Getenv("CONFIG_PATH_TEST")
	}
	return res
}
