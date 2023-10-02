package main

import (
	"fmt"
	"net/http"
)

func getMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello, World!!")
}

func main() {
	http.HandleFunc("/api", getMessage)

	fmt.Println("Server is running on PORT: 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
