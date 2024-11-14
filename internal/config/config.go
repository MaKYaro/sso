package config

import (
	"flag"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Config defines service's configuration
type Config struct {
	Env          string             `yaml:"env" env-default:"prod"`
	DBConnection DBConnectionConfig `yaml:"db_connection" env-required:"true"`
	GRPCServer   GRPCServerConfig   `yaml:"grpc_server" env-required:"true"`
	TokenTTL     time.Duration      `yaml:"token_ttl" env-default:"5m"`
}

// DBConnectionConfig defines db connection configuration:
// user, password, port, host, name
type DBConnectionConfig struct {
	User     string `yaml:"user" env-default:"postgres"`
	Password string `yaml:"password" env-required:"true"`
	Host     string `yaml:"host" env-default:"localhost"`
	Port     int    `yaml:"port" env-default:"5432"`
	DBName   string `yaml:"db_name" env-default:"sso"`
}

// GRPCServerConfig defines grpc server configuration:
// port, address, timeout
type GRPCServerConfig struct {
	Port    int           `yaml:"port" env-default:"8080"`
	Address string        `yaml:"address" env-default:"localhost"`
	Timeout time.Duration `yaml:"timeout" env-default:"3s"`
}

// MustLoad loads configuration from yaml file
// if some error occures or config file path isn't set MustLoad panics
func MustLoad() *Config {
	path := getConfPath()

	if path == "" {
		panic("config path isn't set")
	}

	if _, err := os.Stat(path); err != nil {
		panic("incorrect config file path: " + path)
	}

	configData, err := os.ReadFile(path)
	if err != nil {
		panic("can't read config file: " + err.Error())
	}

	var cfg Config

	err = yaml.Unmarshal(configData, &cfg)
	if err != nil {
		panic("can't parse config file: " + err.Error())
	}

	return &cfg
}

// getConfigPath gets config path from CONFIG_PATH env variable or from flag
// env > flag
func getConfPath() string {
	path := os.Getenv("CONFIG_PATH")
	if path != "" {
		return path
	}

	flag.StringVar(&path, "config-path", "../config/local.yaml", "path to config file")
	flag.Parse()

	return path
}
