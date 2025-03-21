package routes

import (
	"github.com/gorilla/mux"
	"orgmanager/controllers"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/employees", controllers.GetEmployees).Methods("GET")
	r.HandleFunc("/employees/{id}", controllers.GetEmployee).Methods("GET")
	r.HandleFunc("/employees", controllers.CreateEmployee).Methods("POST")
	r.HandleFunc("/employees/{id}", controllers.UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employees/{id}", controllers.DeleteEmployee).Methods("DELETE")

	return r
}
