package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/harranali/task-manager-api/config"
	"github.com/harranali/task-manager-api/internal/task"
	"github.com/harranali/task-manager-api/internal/user"
	"github.com/joho/godotenv"
)

func main() {
	// load env if not production
	if os.Getenv("GO_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			fmt.Println("no .env file found, skipping...")
		}
	}
	// init config
	config.NewConfig()

	const PORT int = 8080
	mux := http.NewServeMux()

	user.RegisterRoutes(mux)
	task.RegisterRoutes(mux)

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", PORT),
		Handler: mux,
	}

	log.Printf("listening on port: %v", PORT)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
