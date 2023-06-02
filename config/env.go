package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetEnv(Key string) string {
	_ = godotenv.Load()
	return os.Getenv(Key)
}

func SetEnvConf(Type string, ConfKey string) {
	ConfSplit := strings.Split(ConfKey, ".")
	var UpConfSplit []string
	for _, v := range ConfSplit {
		UpConfSplit = append(UpConfSplit, strings.ToUpper(v))
	}
	var EnvKey = strings.Join(UpConfSplit, "_")
	env := GetEnv(EnvKey)
	if env != "" {
		switch Type {
		case "string":
			Set(ConfKey, env)
		case "int":
			conv, err := strconv.Atoi(env)
			if err != nil {
				log.Panicf("[Config] Error converting string to int: %v", err)
			}
			Set(ConfKey, conv)
		case "bool":
			switch env {
			case "true":
				Set(ConfKey, true)
			case "false":
				Set(ConfKey, false)
			}
		case "array":
			trim := strings.TrimSpace(env)
			trim = strings.TrimPrefix(trim, "[")
			trim = strings.TrimSuffix(trim, "]")
			var envArray []string
			for _, v := range strings.Split(trim, ",") {
				trimSub := strings.TrimPrefix(v, "\"")
				trimSub = strings.TrimSuffix(trimSub, "\"")
				envArray = append(envArray, trimSub)
			}
			Set(ConfKey, envArray)
		}
	}
}

func EnvInit() {
	var dict = map[string][]string{
		"bool":   {"server.file_logger", "server.debug_mode", "search.enable", "mirai.enable", "logging.external", "logging.internal_log", "cache.external", "gocqhttp.enable", "gocqhttp.enable_ws", "plugin.lua_enable", "plugin.lua_sandbox"},
		"int":    {"search.index_name", "logging.cache_num", "gocqhttp.delay"},
		"string": {"server.listen_port", "server.rpc_port", "search.host", "search.username", "search.password", "logging.mongo_uri", "logging.mongo_db", "cache.master", "cache.username", "cache.password", "database.type", "database.host", "database.name", "database.username", "database.password", "gocqhttp.host_url", "plugin.lua_script_dir"},
		"array":  {"cache.hosts"},
	}
	for key, value := range dict {
		for _, sub := range value {
			SetEnvConf(key, sub)
		}
	}
}
