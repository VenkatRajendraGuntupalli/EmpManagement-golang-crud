package main

import (
	"log"
	"net/http"

	"orgmanager/config"
	"orgmanager/routes"
)

func main() {
	db := config.Connect()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := routes.SetupRouter()
	log.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
