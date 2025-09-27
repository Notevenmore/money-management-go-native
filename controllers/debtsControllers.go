package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"money-management-be/models"
	"money-management-be/responses"
	"money-management-be/services"
)

type DebtControllers struct {
	Service *services.DebtServices
}

func NewDebtController(Service *services.DebtServices) *DebtControllers {
	return &DebtControllers{Service: Service}
}

func (Controller *DebtControllers) IndexHandler(w http.ResponseWriter, r *http.Request) {
	debts, err := Controller.Service.GetAllDebt()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responses.SendSuccess(w, "Daftar Tagihan berhasil didapatkan", debts)
}

func (Controller *DebtControllers) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var debt models.Debt

	if err := json.NewDecoder(r.Body).Decode(&debt); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := Controller.Service.CreateDebt(debt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responses.SendSuccess(w, "Data Tagihan berhasil ditambahkan", nil)
}

func (Controller *DebtControllers) UpdateHandler(w http.ResponseWriter, r *http.Request, id string, db *sql.DB) {
	var debt models.Debt

	if err := json.NewDecoder(r.Body).Decode(&debt); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := Controller.Service.UpdateDebt(debt, id, db); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responses.SendSuccess(w, "Data Tagihan berhasil diupdate", nil)
}
