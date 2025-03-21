package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"orgmanager/config"
	"orgmanager/models"
)

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	var employees []models.Employee
	config.DB.Find(&employees)
	json.NewEncoder(w).Encode(employees)
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var employee models.Employee
	result := config.DB.First(&employee, id)
	if result.Error != nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(employee)
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var employee models.Employee
	json.NewDecoder(r.Body).Decode(&employee)
	config.DB.Create(&employee)
	json.NewEncoder(w).Encode(employee)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var employee models.Employee
	result := config.DB.First(&employee, id)
	if result.Error != nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}
	json.NewDecoder(r.Body).Decode(&employee)
	config.DB.Save(&employee)
	json.NewEncoder(w).Encode(employee)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var employee models.Employee
	result := config.DB.Delete(&employee, id)
	if result.Error != nil {
		http.Error(w, "Failed to delete", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode("Employee deleted successfully")
}
