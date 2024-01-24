// main.go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CrYptOz007/Fusion/internal/database"
	"github.com/CrYptOz007/Fusion/internal/server"
	"github.com/joho/godotenv"
)

var environment string

func init() {
	environment = os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "local"
	}
	if environment == "local" {
		err := godotenv.Load(".env")
		if err != nil {
			os.Exit(1)
		}
		fmt.Println("Running in development mode")
	}
}

func main() {
	fmt.Println("Main: Starting Fusion API Server")

	// Initialise db connection
	connection := new(database.Connection)

	errors := make(chan error)

	go connection.Init(errors)
	fmt.Println("Waiting for database connection to be initialized")
	err := <-errors
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Main: Database connection initialized")

	// Initialise API server
	apiServer := new(server.Server)

	apiServer.Init(connection)
}
