package routes

import (
	"net/http" // Di sini kita akan menggunakan handler yang berada di package main

	"github.com/gorilla/mux"
)

// UserHandlerContract adalah antarmuka yang dibutuhkan SetupRoutes
// agar tidak bergantung langsung pada implementasi *main.UserHandler.
// Anda bisa pindahkan interface ini ke package handler terpisah jika perlu.
type CustomerHandlerContract interface {
	CreateCustomer(w http.ResponseWriter, r *http.Request)
	GetCustomerByID(w http.ResponseWriter, r *http.Request)
	GetAllCustomers(w http.ResponseWriter, r *http.Request)
	UpdateCustomer(w http.ResponseWriter, r *http.Request)
	DeleteCustomer(w http.ResponseWriter, r *http.Request)
}

type FamilyListHandlerInterface interface {
	CreateBulkFamilyList(w http.ResponseWriter, r *http.Request)
	GetFamilyListByID(w http.ResponseWriter, r *http.Request)
	GetAllFamilyListsByUserID(w http.ResponseWriter, r *http.Request)
	UpdateFamilyList(w http.ResponseWriter, r *http.Request)
	DeleteFamilyList(w http.ResponseWriter, r *http.Request)
}

type NationalityHandlerInterface interface {
	GetAllNationalities(w http.ResponseWriter, r *http.Request) // Hanya GET list
}

func SetupRoutes(
	router *mux.Router,
	customerHandler CustomerHandlerContract,
	familyListHandler FamilyListHandlerInterface,
	nationalityHandler NationalityHandlerInterface,
) {

	// Customer Routes (CRUD)
	router.HandleFunc("/customers", customerHandler.CreateCustomer).Methods("POST")
	router.HandleFunc("/customers/{cst_id}", customerHandler.GetCustomerByID).Methods("GET")
	router.HandleFunc("/customers", customerHandler.GetAllCustomers).Methods("GET")
	router.HandleFunc("/customers/{cst_id}", customerHandler.UpdateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{cst_id}", customerHandler.DeleteCustomer).Methods("DELETE")

	// FamilyList Routes (CRUD)
	router.HandleFunc("/family-lists", familyListHandler.CreateBulkFamilyList).Methods("POST")
	router.HandleFunc("/family-lists/{fl_id}", familyListHandler.GetFamilyListByID).Methods("GET")
	router.HandleFunc("/users/{cst_id}/family-lists", familyListHandler.GetAllFamilyListsByUserID).Methods("GET") // Get all family lists for a specific user
	router.HandleFunc("/family-lists/{fl_id}", familyListHandler.UpdateFamilyList).Methods("PUT")
	router.HandleFunc("/family-lists/{fl_id}", familyListHandler.DeleteFamilyList).Methods("DELETE")

	// Nationality Routes (Hanya List)
	router.HandleFunc("/nationalities", nationalityHandler.GetAllNationalities).Methods("GET")
}
