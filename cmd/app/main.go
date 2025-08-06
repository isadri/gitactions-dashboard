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
	}
	if !requiredEnvVarsExist() {
		os.Exit(1)
	}
	server.RegisterFuncs()
	log.Info("starting server")
	log.Info("server listening on 0.0.0.0:8000")
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
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
