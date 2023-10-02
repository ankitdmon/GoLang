package routers

import (
	"GOLANG/controllers"
	"github.com/gorilla/mux"
)

func SetupCRUDRouter(router *mux.Router) {
	router.HandleFunc("/crud/get", controllers.GetMessage).Methods("GET")
	router.HandleFunc("/crud/post", controllers.PostMessage).Methods("POST")
	router.HandleFunc("/crud/put", controllers.PutMessage).Methods("PUT")
	router.HandleFunc("/crud/delete", controllers.DeleteMessage).Methods("DELETE")
}
