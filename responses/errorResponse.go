package responses

import (
	"encoding/json"
	"net/http"

	"money-management-be/utils"
)

func SendError(w http.ResponseWriter, message string, err string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := utils.ErrorResponse{
		Status:  "error",
		Message: message,
		Error:   err,
	}

	json.NewEncoder(w).Encode(response)
}
