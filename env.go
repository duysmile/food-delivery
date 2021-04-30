package main

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	DBConnectionStr string
	S3BucketName    string
	S3Region        string
	S3APIKey        string
	S3Secret        string
	S3Domain        string
	SecretKeyJWT    string
}

func checkEnvFile(file string) error {
	viper.SetConfigFile(file)
	err := viper.ReadInConfig()
	return err
}

func getEnvVar(key string) string {
	return viper.GetString(key)
}

func Init() Env {
	if err := checkEnvFile(".env.yml"); err != nil {
		log.Fatalf("Error read env %s", err)
	}

	var env Env
	env.DBConnectionStr = getEnvVar("DB_CONNECTION")
	env.S3BucketName = getEnvVar("S3_BUCKET_NAME")
	env.S3Region = getEnvVar("S3_REGION")
	env.S3APIKey = getEnvVar("S3_API_KEY")
	env.S3Secret = getEnvVar("S3_SECRET")
	env.S3Domain = getEnvVar("S3_DOMAIN")
	env.SecretKeyJWT = getEnvVar("SECRET_KEY_JWT")
	return env
}
