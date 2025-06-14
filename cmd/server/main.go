package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/harranali/task-manager-api/internal/task"
	"github.com/harranali/task-manager-api/internal/user"
)

func main() {
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
