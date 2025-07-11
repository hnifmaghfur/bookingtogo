package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/hnifmaghfur/bookingtogo/internal/domain"
	"github.com/hnifmaghfur/bookingtogo/internal/interfaces"

	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	customerService interfaces.CustomerService // Dependensi pada antarmuka CustomerService
}

// Pastikan CustomerHandler mengimplementasikan CustomerHandlerContract
var _ CustomerHandlerContract = (*CustomerHandler)(nil)

func NewCustomerHandler(customerService interfaces.CustomerService) *CustomerHandler {
	return &CustomerHandler{customerService: customerService}
}

func (h *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer domain.Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.customerService.CreateCustomer(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)
}

func (h *CustomerHandler) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["cst_id"], 10, 64)
	if err != nil {
		http.Error(w, "ID pengguna tidak valid", http.StatusBadRequest)
		return
	}

	customer, err := h.customerService.GetCustomerByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

func (h *CustomerHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.customerService.GetAllCustomers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (h *CustomerHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["cst_id"], 10, 64)
	if err != nil {
		http.Error(w, "ID pengguna tidak valid", http.StatusBadRequest)
		return
	}

	var customer domain.Customer
	err = json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	customer.CstID = uint(id) // Pastikan ID dari URL digunakan

	err = h.customerService.UpdateCustomer(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

func (h *CustomerHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["cst_id"], 10, 64)
	if err != nil {
		http.Error(w, "ID pengguna tidak valid", http.StatusBadRequest)
		return
	}

	err = h.customerService.DeleteCustomer(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
