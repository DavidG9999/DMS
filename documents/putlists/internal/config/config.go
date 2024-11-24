package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env  string     `yaml:"env" env-default:"local"`
	DB   DBConfig   `yaml:"db" env-required:"true"`
	GRPC GRPCConfig `yaml:"grpc"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}
	return MustLoadByPath(configPath)
}

func MustLoadByPath(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + err.Error())
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
