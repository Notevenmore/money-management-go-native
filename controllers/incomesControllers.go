package controllers

import (
	"encoding/json"
	"net/http"

	"money-management-be/models"
	"money-management-be/responses"
	"money-management-be/services"
)

type IncomeControllers struct {
	Service *services.IncomeServices
}

func NewIncomeController(Service *services.IncomeServices) *IncomeControllers {
	return &IncomeControllers{Service: Service}
}

func (Controller *IncomeControllers) IndexHandler(w http.ResponseWriter, r *http.Request) {
	income, err := Controller.Service.GetAllIncomes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responses.SendSuccess(w, "Daftar Pemasukan berhasil didapatkan", income)
}

func (Controller *IncomeControllers) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var income models.Income
	if err := json.NewDecoder(r.Body).Decode(&income); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := Controller.Service.CreateIncome(income); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responses.SendSuccess(w, "Data Pemasukan berhasil ditambahkan", nil)

}

// func UpdateHandler(w http.ResponseWriter, r *http.Request, id string) {
// 	response := map[string]interface{}{
// 		"message": "Hello Update Handler Incomes",
// 		"status":  true,
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(response)
// }

// func DeleteHandler(w http.ResponseWriter, r *http.Request, id string) {
// 	response := map[string]interface{}{
// 		"message": "Hello Delete Handler Incomes",
// 		"status":  true,
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(response)
// }
