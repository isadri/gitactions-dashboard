package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/isadri/cicd-dashboard/internal/gitactions"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var workflows *gitactions.Workflows

	if workflows, err = gitactions.GetWorkflows("isadri", "eval"); err != nil {
		log.Fatal(err)
	}
	data, err := json.MarshalIndent(*workflows, "", "    ")
	if err != nil {
		log.Fatalf("error marshaling JSON: %s", err)
	}
	fmt.Println(string(data))
}
