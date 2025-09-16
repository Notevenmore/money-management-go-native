package assets

import (
	"net/http"
	"strings"

	"money-management-be/controllers/assetsControllers"
)

func InitAssetsRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/")

		if path == "" {
			switch r.Method {
			case http.MethodGet:
				assetsControllers.IndexHandler(w, r)
			case http.MethodPost:
				assetsControllers.CreateHandler(w, r)
			default:
				http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
			}

			return
		}

		id := path
		switch r.Method {
		case http.MethodPut:
			assetsControllers.UpdateHandler(w, r, id)
		default:
			http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
