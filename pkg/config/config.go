package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"

	_ "github.com/lib/pq"
)

type Config struct {
	HTTP_PORT          string
	HTTP_USER_POST     string
	GRPC_USER_PORT     string
	GRPC_LEARNING_PORT string
	GRPC_PROGRESS_PORT string
	DB_HOST            string
	DB_PORT            string
	DB_USER            string
	DB_NAME            string
	DB_PASSWORD        string
	DB_CASBIN_DRIVER   string
	ACCESS_TOKEN       string
	REFRESH_TOKEN      string
}

func Load() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found")
	}

	config := Config{}

	config.HTTP_PORT = cast.ToString(coalesce("HTTP_PORT", "8080"))
	config.HTTP_USER_POST = cast.ToString(coalesce("HTTP_USER_POST", "8081"))
	config.GRPC_USER_PORT = cast.ToString(coalesce("GRPC_USER_PORT", "50050"))
	config.GRPC_LEARNING_PORT = cast.ToString(coalesce("GRPC_LEARNING_PORT", "50051"))
	config.GRPC_PROGRESS_PORT = cast.ToString(coalesce("GRPC_PROGRESS_PORT", "50052"))
	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "localhost"))
	config.DB_PORT = cast.ToString(coalesce("DB_PORT", "5432"))
	config.DB_NAME = cast.ToString(coalesce("DB_NAME", "ecommerce_auth_service"))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "postgres"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "123321"))
	config.DB_CASBIN_DRIVER = cast.ToString(coalesce("DB_CASBIN_DRIVER", "postgres"))
	config.ACCESS_TOKEN = cast.ToString(coalesce("ACCESS_TOKEN", "key_is_really_easy"))
	config.REFRESH_TOKEN = cast.ToString(coalesce("REFRESH_TOKEN", "key_is_not_hard"))

	return config
}

func coalesce(env string, defaultValue interface{}) interface{} {
	value, exists := os.LookupEnv(env)
	if !exists {
		return defaultValue
	}
	return value
}
