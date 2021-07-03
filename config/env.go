package config

import (
	"os"
	"time"
)

var (
	ConnectTime	  = GetEnv("CONNECT_TIME", time.Second * time.Duration(5))
	ServerPort    = GetEnv("SERVER_PORT", "8080")
	MongoUrl      = GetEnv("MONGODB_URL", "mongodb://127.0.0.1:27017/?connect=direct")
)

func GetEnv(key string, defaultValue interface{}) interface{} {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}