package debtsControllers

import (
	"encoding/json"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"message": "Hello Index Handler Debts",
		"status":  true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"message": "Hello Create Handler Debts",
		"status":  true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request, id string) {
	response := map[string]interface{}{
		"message": "Hello Update Handler Debts",
		"status":  true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
