package main

import (
	"errors"
	"os"
)

func GetOsEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func SetOsEnv(env string, key string) {
	os.Setenv(env, key)
}

func GetServer() (string, error) {
	server := GetOsEnv("EXCHANGE_SERVER", "")
	if len(server) == 0 {
		return "", errors.New("You need to set the environment EXCHANGE_SERVER (ex: localhost)")
	}
	return server, nil
}
