package responses

import (
	"encoding/json"
	"net/http"

	"money-management-be/utils"
)

func SendSuccess(w http.ResponseWriter, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := utils.SuccessResponse{
		Status:     "Success",
		Message:    message,
		Data:       data,
		StatusCode: http.StatusAccepted,
	}

	json.NewEncoder(w).Encode(response)
}
