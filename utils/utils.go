package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type ErrorResponse struct {
	Status  string `json:"status"`
	Message any    `json:"message"`
}

type SuccessResponse struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

func WriteSuccessResponse(w http.ResponseWriter, code int, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	response, _ := json.Marshal(SuccessResponse{
		Status: "success",
		Data:   data,
	})
	fmt.Fprint(w, string(response))
}

func WriteErrorResponse(w http.ResponseWriter, code int, message any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	response, _ := json.Marshal(ErrorResponse{
		Status:  "error",
		Message: message,
	})
	fmt.Fprint(w, string(response))
}

func GetEnvMust(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		log.Fatalf("unable to get env var: %v", key)
	}
	return value
}

func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
