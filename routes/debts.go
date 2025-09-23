package routes

import (
	"net/http"
	"strings"

	"money-management-be/controllers"
)

func InitDebtsRoutes(Controller *controllers.DebtControllers) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/")

		if path == "" {
			switch r.Method {
			case http.MethodGet:
				Controller.IndexHandler(w, r)
			case http.MethodPost:
				Controller.CreateHandler(w, r)
			default:
				http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
			}

			return
		}

		// id := path
		// switch r.Method {
		// case http.MethodPut:
		// 	debtsControllers.UpdateHandler(w, r, id)
		// default:
		// 	http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		// }
	})

	return mux
}
