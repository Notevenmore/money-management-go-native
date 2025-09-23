package controllers

import (
	"encoding/json"
	"net/http"

	"money-management-be/models"
	"money-management-be/responses"
	"money-management-be/services"
)

type OutcomeControllers struct {
	Service *services.OutcomeServices
}

func NewOutcomeController(Service *services.OutcomeServices) *OutcomeControllers {
	return &OutcomeControllers{Service: Service}
}

func (Controller *OutcomeControllers) IndexHandler(w http.ResponseWriter, r *http.Request) {
	outcomes, err := Controller.Service.GetAllOutcomes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responses.SendSuccess(w, "Daftar Pengeluaran berhasil didapatkan", outcomes)
}

func (Controller *OutcomeControllers) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var outcome models.Outcome
	if err := json.NewDecoder(r.Body).Decode(&outcome); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := Controller.Service.CreateOutcome(outcome); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responses.SendSuccess(w, "Data Pengeluaran berhasil ditambahkan", nil)
}

// func UpdateHandler(w http.ResponseWriter, r *http.Request, id string) {
// 	response := map[string]interface{}{
// 		"message": "Hello Update Handler Outcomes",
// 		"status":  true,
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(response)
// }

// func DeleteHandler(w http.ResponseWriter, r *http.Request, id string) {
// 	response := map[string]interface{}{
// 		"message": "Hello Delete Handler Outcomes",
// 		"status":  true,
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(response)
// }
