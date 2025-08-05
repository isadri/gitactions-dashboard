package main

import (
	"log"
	"net/http"

	"github.com/isadri/cicd-dashboard/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	log.Print("loading environment variables in .env")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	server.RegisterFuncs()
	log.Print("starting server")
	log.Print("server listening on localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
