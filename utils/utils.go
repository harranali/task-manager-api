package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
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

func WriteErrorResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	response, _ := json.Marshal(ErrorResponse{
		Status:  "error",
		Message: message,
	})
	fmt.Fprint(w, string(response))
}
