package config

import (
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
			log.Fatalf("[Config] Error writing config: %v", err)
		}
	}
	if err := viper.Unmarshal(conf); err != nil {
		log.Fatal(err)
	}
}

func Get(key string) any {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("[Config] Error reading config: %v", err)
	}
	return viper.Get(key)
}

func Set(key string, value any) {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("[Config] Error reading config: %v", err)
	}
	viper.Set(key, value)
}
