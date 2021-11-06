package handleRequests

import (
	"log"
	"net/http"

	controller "github.com/brandonhsz/golangApi/controllers"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", controller.IndexRoute).Methods("GET")
	router.HandleFunc("/tasks", controller.GetTask).Methods("GET")
	router.HandleFunc("/tasks", controller.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", controller.GetAnyTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", controller.DeleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}
