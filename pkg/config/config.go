package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type MySQL struct {
	DSN string `mapstructure:"dsn"`
}

type Redis struct {
	Host string `mapstructure:"host"`
	Port int32 `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB int `mapstructure:"db"`
}

// Config is a structure that holds the application's configuration.
type Config struct {
	MySQL `mapstructure:"mysql"`
	Redis `mapstructure:"redis"`
}

// ConfigData holds the parsed configuration.
var ConfigData *Config

// InitConfig initializes the configuration using Viper and binds it to the Config structure.
func InitConfig(configFile string) {
	viper.SetConfigName(configFile)           // name of config file (without extension)
	viper.AddConfigPath("./config/")              // path to look for the config file in
	viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
	viper.SetEnvPrefix("ming")          // prefix to use when reading environment variables
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // replace dots with underscores in env vars

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Unable to read config file, %v", err)
	}

	// Bind configuration data to the Config structure
	var conf Config
	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	ConfigData = &conf

	// Set environment variable overrides
	viper.AutomaticEnv()
}

// GetConfig returns the initialized Config structure.
func GetConfig() *Config {
	return ConfigData
}