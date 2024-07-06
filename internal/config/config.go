package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/spf13/viper"
	"log"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {
	viper.SetConfigName("local")    // имя файла конфигурации (без расширения)
	viper.SetConfigType("yaml")     // тип файла конфигурации
	viper.AddConfigPath("./config") // путь к файлу конфигурации

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(viper.ConfigFileUsed(), &cfg); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	return &cfg
}
