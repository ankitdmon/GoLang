package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!!")
}

func main() {
	router := mux.NewRouter()

	// Define the route using the Gorilla Mux router
	router.HandleFunc("/test", getMessage).Methods("GET")

	fmt.Println("Server is running on PORT: 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
