package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App       AppConfig       `mapstructure:"app"`
	Database  DatabaseConfig  `mapstructure:"database"`
	CORS      CORSConfig      `mapstructure:"cors"`
	Cache     CacheConfig     `mapstructure:"cache"`
	RateLimit RateLimitConfig `mapstructure:"rate_limit"`
	Storage   StorageConfig   `mapstructure:"storage"`
	Logging   LoggingConfig   `mapstructure:"logging"`
}

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Env     string `mapstructure:"env"`
	Port    int    `mapstructure:"port"`
	Version string `mapstructure:"version"`
}

type DatabaseConfig struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	Name            string `mapstructure:"name"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
}

type CORSConfig struct {
	AllowedOrigins []string `mapstructure:"allowed_origins"`
	AllowedMethods []string `mapstructure:"allowed_methods"`
	AllowedHeaders []string `mapstructure:"allowed_headers"`
}

type CacheConfig struct {
	Enabled bool        `mapstructure:"enabled"`
	Type    string      `mapstructure:"type"`
	Redis   RedisConfig `mapstructure:"redis"`
	TTL     int         `mapstructure:"ttl"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type RateLimitConfig struct {
	Enabled  bool `mapstructure:"enabled"`
	Requests int  `mapstructure:"requests"`
	Window   int  `mapstructure:"window"`
}

type StorageConfig struct {
	CDNURL string `mapstructure:"cdn_url"`
}

type LoggingConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file: %v", err)
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Printf("Error unmarshaling config: %v", err)
		return nil, err
	}

	return &config, nil
}
