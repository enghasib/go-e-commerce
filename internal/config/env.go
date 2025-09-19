package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type env struct {
	Version     float64
	ServiceName string
	HttpPort    uint
	JwtSecret   string
}

var Env *env

func LoadEnv() *env {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error to load env!")
		os.Exit(1)
	}

	serviceName, err := getEnv("SERVICE_NAME")
	if err != nil {
		log.Fatal("Service name", err)
		os.Exit(1)
	}

	version, err := getEnv("VERSION")
	if err != nil {
		log.Fatal("Version", err)
		os.Exit(1)
	}
	versionFloat64, _ := strconv.ParseFloat(version, 64)

	port, err := getEnv("HTTP_PORT")
	if err != nil {
		log.Fatal("Port", err)
		os.Exit(1)
	}
	httpPort, _ := strconv.Atoi(port)

	jwtSecret, err := getEnv("JWT_SECRET")
	if err != nil {
		log.Fatal("JWT secret", err)
		os.Exit(1)
	}

	Env = &env{
		ServiceName: serviceName,
		Version:     versionFloat64,
		HttpPort:    uint(httpPort),
		JwtSecret:   jwtSecret,
	}
	return Env
}

func getEnv(key string) (string, error) {
	if value, ok := os.LookupEnv(key); ok {
		return value, nil
	}
	return "", errors.New(" env can't load")
}
