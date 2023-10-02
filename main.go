package main

import (
	"fmt"
	"net/http"

	"GOLANG/routers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	routers.SetupCRUDRouter(router)

	fmt.Println("Server is running on PORT: 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
