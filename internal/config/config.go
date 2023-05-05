package config

import (
	"github.com/Elyart-Network/NyaBot/internal/logger"
	"github.com/spf13/viper"
	"log"
)

func init() {
	// Read global config
	viper.SetConfigFile("config.yaml")
	viper.SetDefault("Server", serverDef)
	viper.SetDefault("Mirai", miraiDef)
	viper.SetDefault("GoCqHttp", goCqHttpDef)
	viper.SetDefault("Database", databaseDef)
	viper.SetDefault("Logging", loggingDef)
	viper.SetDefault("Cache", cacheDef)
	viper.SetDefault("Search", searchDef)
	viper.SetDefault("Plugin", pluginDef)
	conf := &config{}
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Can't read config, trying to modify!")
		if err := viper.WriteConfig(); err != nil {
			logger.Errorf("Config", "Error writing config", err.Error())
		}
	}
	if err := viper.Unmarshal(conf); err != nil {
		logger.ErrorFatal(err)
	}
}

func Get(key string) interface{} {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		logger.Errorf("Config", "Error reading config", err.Error())
	}
	return viper.Get(key)
}

func Set(key string, value interface{}) {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		logger.Errorf("Config", "Error reading config", err.Error())
	}
	viper.Set(key, value)
}
