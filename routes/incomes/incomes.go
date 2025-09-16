package incomes

import (
	"net/http"
	"strings"

	"money-management-be/controllers/incomesControllers"
)

func InitIncomesRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/")

		if path == "" {
			switch r.Method {
			case http.MethodGet:
				incomesControllers.IndexHandler(w, r)
			case http.MethodPost:
				incomesControllers.CreateHandler(w, r)
			default:
				http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
			}

			return
		}

		id := path
		switch r.Method {
		case http.MethodPut:
			incomesControllers.UpdateHandler(w, r, id)
		case http.MethodDelete:
			incomesControllers.DeleteHandler(w, r, id)
		default:
			http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
