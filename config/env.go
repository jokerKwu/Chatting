package config

import (
	"os"
)

var (
	ConnectTime	  = GetEnv("CONNECT_TIME", 5)
	ServerPort    = GetEnv("SERVER_PORT", "8080")
	MongoUrl      = GetEnv("MONGODB_URL", "mongodb://127.0.0.1:27017/?connect=direct")
	AcceptTokenExp = GetEnv("ACCEPT_TOKEN_EXP", 1)
	RefreshTokenExp = GetEnv("REFRESH_TOKEN_EXP",24 * 7)
)

func GetEnv(key string, defaultValue interface{}) interface{} {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}