package main

import (
	"net/http"
	"os"

	"github.com/isadri/gitactions-dashboard/internal/server"
	"github.com/isadri/gitactions-dashboard/internal/utils"
	"github.com/joho/godotenv"
)

func main() {
	log := utils.GetLogger()

	log.Info("loading environment variables in .env")
	if envFileExists() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	} else {
		log.Info(".env does not exist. Read from system environment variables")
	}
	if !requiredEnvVarsExist() {
		os.Exit(1)
	}
	server.RegisterFuncs()
	address := getBindAddress()
	log.Info("starting server")
	log.Infof("server listening on %s", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

func getBindAddress() string {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}
	ip := os.Getenv("APP_BIND")
	if ip == "" {
		ip = "localhost"
	}
	return ip + ":" + port
}

func envFileExists() bool {
	info, err := os.Stat(".env")
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func requiredEnvVarsExist() bool {
	log := utils.GetLogger()

	if os.Getenv("GITHUB_TOKEN") == "" {
		log.Error("missing GITHUB_TOKEN environment variable")
		return false
	}
	if os.Getenv("ORG_NAME") == "" {
		log.Error("missing ORG_NAME environment variable")
		return false
	}
	return true
}
