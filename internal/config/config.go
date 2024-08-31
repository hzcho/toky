package config

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env       string        `yaml:"env" env-required:"true"`
	TokenTTL  time.Duration `yaml:"token_ttl" env-required:"true"`
	UploadDir string        `yaml:"upload_dir" env-required:"true"`
	DB        DB            `yaml:"db" env-required:"true"`
	Server    Server        `yaml:"server" env-required:"true"`
}

type DB struct {
	Username string `yaml:"username" env-required:"true"`
	Host     string `yaml:"host" env-required:"true"`
	Port     string `yaml:"port" env-required:"true"`
	DBName   string `yaml:"dbname" env-required:"true"`
	SSLMode  string `yaml:"sslmode" env-required:"true"`
	Password string
}

type Server struct {
	Port    string        `yaml:"port" env-required:"true"`
	Timeout time.Duration `yaml:"timeout" env-required:"true"`
}

func MustLoad(path string) *Config {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	if path == "" {
		panic("path is empty")
	}

	cfg := Config{}
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic(err)
	}

	cfg.DB.Password = os.Getenv("DB_PASSWORD")
	if cfg.DB.Password == "" {
		panic("password of database is empty")
	}

	return &cfg
}
