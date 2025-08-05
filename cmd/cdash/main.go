package main

import (
	"net/http"

	"github.com/isadri/cicd-dashboard/internal/server"
	"github.com/isadri/cicd-dashboard/internal/utils"
	"github.com/joho/godotenv"
)

func main() {
	log := utils.GetLogger()

	log.Info("loading environment variables in .env")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	server.RegisterFuncs()
	log.Info("starting server")
	log.Info("server listening on localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
