package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	config configData
}

type configData struct {
	// http
	HTTPPort int    `mapstructure:"http_port"`
	GinMode  string `mapstructure:"gin_mode"`

	// log
	LogLevel        logrus.Level `mapstructure:"log_level"`
	LogFileLocation string       `mapstructure:"log_file_location"`

	// database
	DBType           string `mapstructure:"db_type"`
	DBConnectionPath string `mapstructure:"db_path"`
}

func NewConfig() (*Config, error) {

	viper.SetConfigName("config") // name of config.yaml file (without extension)
	viper.SetConfigType("yaml")

	// where to look for
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config.yaml file
	if err != nil {             // Handle errors reading the config.yaml file
		return nil, err
	}

	viper.AutomaticEnv()

	config := configData{}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &Config{config: config}, nil
}

func (c *Config) HTTPPort() int {
	return c.config.HTTPPort
}

func (c *Config) DBConnectionPath() string {
	return c.config.DBConnectionPath
}

func (c *Config) GinMode() string {
	return c.config.GinMode
}

func (c *Config) LogFileLocation() string {
	return c.config.LogFileLocation
}

func (c *Config) LogLevel() logrus.Level {
	return c.config.LogLevel
}
