package main

import (
	"cacheDataService/handlers"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/api/set", handlers.SetTableCache)
	http.HandleFunc("/api/get", handlers.GetTableCache)

	fmt.Printf("The server started...")
	http.ListenAndServe(":8090", nil)
}
