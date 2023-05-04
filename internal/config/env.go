package config

import (
	"github.com/Elyart-Network/NyaBot/internal/logger"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

func GetEnv(Key string) string {
	_ = godotenv.Load()
	return os.Getenv(Key)
}

func SetEnvConf(Type string, ConfKey string, EnvKey string) {
	env := GetEnv(EnvKey)
	if env != "" {
		switch Type {
		case "string":
			Set(ConfKey, env)
		case "int":
			conv, err := strconv.Atoi(env)
			if err != nil {
				logger.Errorf("Config", "Error converting string to int", err)
			}
			Set(ConfKey, conv)
		case "bool":
			switch env {
			case "true":
				Set(ConfKey, true)
			case "false":
				Set(ConfKey, false)
			}
		}
	}
}

func EnvInit() {
	SetEnvConf("string", "server.listen_port", "SERVER_LISTEN_PORT")
	SetEnvConf("bool", "server.file_logger", "SERVER_FILE_LOGGER")
	SetEnvConf("bool", "server.debug_mode", "SERVER_DEBUG_MODE")
	SetEnvConf("bool", "search.enable", "SEARCH_ENABLED")
	SetEnvConf("string", "search.host", "SEARCH_HOST")
	SetEnvConf("int", "search.index_name", "SEARCH_INDEX_NAME")
	SetEnvConf("string", "search.username", "SEARCH_USERNAME")
	SetEnvConf("string", "search.password", "SEARCH_PASSWORD")
	SetEnvConf("bool", "mirai.enable", "MIRAI_ENABLE")
	SetEnvConf("bool", "logging.external", "LOGGING_EXTERNAL")
	SetEnvConf("string", "logging.mongo_uri", "LOGGING_MONGO_URI")
	SetEnvConf("string", "logging.username", "LOGGING_USERNAME")
	SetEnvConf("string", "logging.password", "LOGGING_PASSWORD")
	SetEnvConf("int", "logging.cache_num", "LOGGING_CACHE_NUM")
	SetEnvConf("bool", "logging.internal_log", "LOGGING_INTERNAL_LOG")
	SetEnvConf("bool", "cache.external", "CACHE_EXTERNAL")
	SetEnvConf("string", "cache.host", "CACHE_HOST")
	SetEnvConf("string", "cache.index_name", "CACHE_INDEX_NAME")
	SetEnvConf("string", "cache.username", "CACHE_USERNAME")
	SetEnvConf("string", "cache.password", "CACHE_PASSWORD")
	SetEnvConf("string", "database.type", "DATABASE_TYPE")
	SetEnvConf("string", "database.host", "DATABASE_HOST")
	SetEnvConf("string", "database.name", "DATABASE_NAME")
	SetEnvConf("string", "database.username", "DATABASE_USERNAME")
	SetEnvConf("string", "database.password", "DATABASE_PASSWORD")
	SetEnvConf("bool", "gocqhttp.enable", "GOCQHTTP_ENABLED")
	SetEnvConf("string", "gocqhttp.host_url", "GOCQHTTP_HOST_URL")
	SetEnvConf("int", "gocqhttp.delay", "GOCQHTTP_DELAY")
	SetEnvConf("bool", "gocqhttp.enable_ws", "GOCQHTTP_ENABLE_WS")
}
