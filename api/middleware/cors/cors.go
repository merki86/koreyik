package cors

import (
	"log/slog"
	"net/http"
)

func New(log *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		log.Info("CORS logger middleware initialized")

		fn := func(w http.ResponseWriter, r *http.Request) {
			// Allow all origins
			w.Header().Set("Access-Control-Allow-Origin", "*")

			// Allow all methods
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

			// Allow common headers
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			// Handle preflight requests
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			// Pass to next handler
			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
