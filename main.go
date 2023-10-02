package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get Method!!")
}

func postMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Post Method!!")
}

func putMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Put Method!!")
}

func deleteMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Delete Method!!")
}

func main() {
	router := mux.NewRouter()

	// Define the route using the Gorilla Mux router
	router.HandleFunc("/get", getMessage).Methods("GET")
	router.HandleFunc("/post", postMessage).Methods("POST")
	router.HandleFunc("/put", putMessage).Methods("PUT")
	router.HandleFunc("/delete", deleteMessage).Methods("DELETE")

	fmt.Println("Server is running on PORT: 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
