package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received POST reques!")

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	fmt.Println("Request Body:", string(body))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", postHandler).Methods("POST")
	fmt.Println("Server is running on http://localhost:8080")
	// Define your routes here
	log.Fatal(http.ListenAndServe(":8080", r))
}
