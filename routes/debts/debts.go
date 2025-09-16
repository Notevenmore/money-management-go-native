package debts

import (
	"net/http"
	"strings"

	"money-management-be/controllers/debtsControllers"
)

func InitDebtsRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/")

		if path == "" {
			switch r.Method {
			case http.MethodGet:
				debtsControllers.IndexHandler(w, r)
			case http.MethodPost:
				debtsControllers.CreateHandler(w, r)
			default:
				http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
			}

			return
		}

		id := path
		switch r.Method {
		case http.MethodPut:
			debtsControllers.UpdateHandler(w, r, id)
		default:
			http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
