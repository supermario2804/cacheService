package main

import (
	"cacheDataService/handlers"
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/api/set", handlers.SetTableCache)
	http.HandleFunc("/api/get", headers)

	fmt.Printf("The server started...")
	http.ListenAndServe(":8090", nil)
}
