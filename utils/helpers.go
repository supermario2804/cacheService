package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func SendHTTPResponse(w http.ResponseWriter, model interface{}) {

	json, err := json.Marshal(model)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusRequestTimeout)
		Println(w, "Something went wrong", http.StatusRequestTimeout)
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, string(json))
}


type ApiResponse struct {
	Success bool                   `json:"success"`
	Status  int                    `json:"statusCode"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
	Errors  map[string]string      `json:"errors,omitempty"`
}
