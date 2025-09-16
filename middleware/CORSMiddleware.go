package middleware

import (
	"net/http"
	"os"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		appOrigin := os.Getenv("APP_ORIGIN")
		appServer := os.Getenv("APP_SERVER")
		if appOrigin == "" || appServer == "" {
			appOrigin = "*"
		} else {
			appOrigin = appServer + appOrigin
		}

		w.Header().Set("Access-Control-Allow-Origin", appOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
