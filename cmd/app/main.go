package main

import (
	"net/http"
	"os"

	"github.com/isadri/gitactions-dashboard/internal/server"
	"github.com/isadri/gitactions-dashboard/internal/utils"
	"github.com/joho/godotenv"
)

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

func main() {
	log := utils.GetLogger()

	log.Info("loading environment variables in .env")
	err := godotenv.Load()
	if !requiredEnvVarsExist() {
		os.Exit(1)
	}
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	server.RegisterFuncs()
	log.Info("starting server")
	log.Info("server listening on localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
