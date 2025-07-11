package routes

import (
	"encoding/json"
	"net/http"

	"github.com/hnifmaghfur/bookingtogo/internal/interfaces"
)

type NationalityHandler struct {
	nationalityService interfaces.NationalityService
}

func NewNationalityHandler(nationalityService interfaces.NationalityService) *NationalityHandler {
	return &NationalityHandler{nationalityService: nationalityService}
}

func (h *NationalityHandler) GetAllNationalities(w http.ResponseWriter, r *http.Request) {
	nationalities, err := h.nationalityService.GetAllNationalities()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nationalities)
}
