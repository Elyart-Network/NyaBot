package config

import (
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.SetConfigFile("configs/config.yaml")
	viper.SetDefault("Server", serverDef)
	viper.SetDefault("GoCqHttp", goCqHttpDef)
	viper.SetDefault("Database", databaseDef)
	viper.SetDefault("Queue", queueDef)
	viper.SetDefault("Cache", cacheDef)
	viper.SetDefault("Search", searchDef)
	conf := &config{}
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Can't read config, trying to modify!")
		if err := viper.WriteConfig(); err != nil {
			log.Panicln("Error writing config!")
		}
	}
	if err := viper.Unmarshal(conf); err != nil {
		log.Fatal(err)
	}
}

func Get(key string) interface{} {
	if err := viper.ReadInConfig(); err != nil {
		log.Panicln("Error reading config!")
	}
	return viper.Get(key)
}
