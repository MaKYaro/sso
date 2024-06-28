package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"
)

type Config struct {
	Env      string        `json:"Env" env-default:"local"`
	TokenTTL time.Duration `json:"TokenTTL" env-required:"true"`
	DBConn   DBConnConfig  `json:"DBConn" env-required:"true"`
	GRPC     GRPCConfig    `json:"GRPC" env-required:"true"`
}

type DBConnConfig struct {
	Host     string `json:"Host"`
	Port     int    `json:"Port"`
	User     string `json:"User"`
	Password string `json:"Password"`
	DBName   string `json:"DBName"`
}

type GRPCConfig struct {
	Port    int           `json:"Port"`
	Timeout time.Duration `json:"Timeout"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	var cfg Config
	if err := readConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

func readConfig(path string, cfg *Config) error {
	const op = "internal.config.readConfig"

	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

// Priority: flag > env > default
// Default value is empty string
func fetchConfigPath() string {
	var res string

	// --config="./config/local.json"
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
