package main

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	DBConnectionStr string
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
	env.DBConnectionStr = getEnvVar("DBConnectionStr")
	return env
}
