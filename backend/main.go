// main.go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CrYptOz007/Fusion/internal/database"
	"github.com/CrYptOz007/Fusion/internal/helpers"
	"github.com/CrYptOz007/Fusion/internal/server"
)

var environment string

func init() {
	environment = os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "development"
		fmt.Println("Running in development mode")
	}
}

func main() {
	fmt.Println("Main: Starting Fusion API Server")

	// Initialise db connection
	connection := new(database.Connection)

	errors := make(chan error)

	errs := os.Setenv("AUTH_KEY", helpers.GenerateRandomKey())
	if errs != nil {
		log.Fatal("Failed to generate auth key")
		os.Exit(1)
	}
	errs = os.Setenv("REFRESH_KEY", helpers.GenerateRandomKey())
	if errs != nil {
		log.Fatal("Failed to generate refresh key")
		os.Exit(1)
	}

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
