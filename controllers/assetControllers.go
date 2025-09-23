package controllers

import (
	"encoding/json"
	"net/http"

	"money-management-be/models"
	"money-management-be/responses"
	"money-management-be/services"
)

type AssetControllers struct {
	Service *services.AssetService
}

func NewAssetController(Service *services.AssetService) *AssetControllers {
	return &AssetControllers{Service: Service}
}

func (Controller *AssetControllers) IndexHandler(w http.ResponseWriter, r *http.Request) {
	assets, err := Controller.Service.GetAllAssets()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responses.SendSuccess(w, "Daftar Aset berhasil didapatkan", assets)
}

func (Controller *AssetControllers) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var asset models.Assets
	if err := json.NewDecoder(r.Body).Decode(&asset); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := Controller.Service.CreateAsset(asset); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responses.SendSuccess(w, "Data Aset berhasil ditambahkan", nil)
}

// func UpdateHandler(w http.ResponseWriter, r *http.Request, id string) {
// 	response := map[string]interface{}{
// 		"message": "Hello Update Handler Aset",
// 		"status":  true,
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(response)
// }
