package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/hnifmaghfur/bookingtogo/internal/domain"
	"github.com/hnifmaghfur/bookingtogo/internal/interfaces"

	"github.com/gorilla/mux"
)

type FamilyListHandler struct {
	familyListService interfaces.FamilyListService
}

func NewFamilyListHandler(familyListService interfaces.FamilyListService) *FamilyListHandler {
	return &FamilyListHandler{familyListService: familyListService}
}

func (h *FamilyListHandler) CreateBulkFamilyList(w http.ResponseWriter, r *http.Request) {
	var familyLists []domain.FamilyList
	if err := json.NewDecoder(r.Body).Decode(&familyLists); err != nil {
		http.Error(w, "Invalid JSON array: "+err.Error(), http.StatusBadRequest)
		return
	}
	if len(familyLists) == 0 {
		http.Error(w, "Family list array cannot be empty", http.StatusBadRequest)
		return
	}
	cstID := familyLists[0].CstID
	for _, f := range familyLists {
		if f.CstID != cstID {
			http.Error(w, "All family members must have the same cst_id", http.StatusBadRequest)
			return
		}
	}
	if err := h.familyListService.CreateBulkFamilyList(familyLists); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(familyLists)
}

func (h *FamilyListHandler) GetFamilyListByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["fl_id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid family list ID", http.StatusBadRequest)
		return
	}

	familyList, err := h.familyListService.GetFamilyListByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(familyList)
}

func (h *FamilyListHandler) GetAllFamilyListsByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.ParseUint(vars["cst_id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	familyLists, err := h.familyListService.GetAllFamilyListsByUserID(uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(familyLists)
}

func (h *FamilyListHandler) UpdateFamilyList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["fl_id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid family list ID", http.StatusBadRequest)
		return
	}

	var familyList domain.FamilyList
	if err := json.NewDecoder(r.Body).Decode(&familyList); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	familyList.FlID = uint(id) // Pastikan ID dari URL digunakan

	if err := h.familyListService.UpdateFamilyList(&familyList); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(familyList)
}

func (h *FamilyListHandler) DeleteFamilyList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["fl_id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid family list ID", http.StatusBadRequest)
		return
	}

	if err := h.familyListService.DeleteFamilyList(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
