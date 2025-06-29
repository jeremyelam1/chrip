package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

func main() {
	mux := http.NewServeMux()

	apiConfig := NewAPIConfig()

	mux.Handle(
		"GET /app/",
		apiConfig.middlewareMetricsInc(
			http.StripPrefix("/app/", http.FileServer(http.Dir(".")))),
	)
	mux.Handle(
		"GET /api/assets",
		apiConfig.middlewareMetricsInc(http.StripPrefix("/api/app/", http.FileServer(http.Dir("./assets/logo.png")))),
	)

	mux.HandleFunc("GET /admin/metrics", apiConfig.adminMetrics)

	mux.HandleFunc("GET /api/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	mux.HandleFunc("GET /api/metrics", apiConfig.handlerMetrics)

	mux.HandleFunc("POST /admin/reset", apiConfig.handlerReset)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Server started on :8080")
	server.ListenAndServe()
}

type apiConfig struct {
	fileserverHits atomic.Int32
}

func NewAPIConfig() *apiConfig {
	return &apiConfig{}
}

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hits: %d", cfg.fileserverHits.Load())))
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	cfg.fileserverHits.Store(0)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset to 0"))
}

func (cfg *apiConfig) adminMetrics(w http.ResponseWriter, r *http.Request) {
	// Format the content with the hits value
	formattedContent := fmt.Sprintf(`<html>
  <body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
  </body>
</html>`, cfg.fileserverHits.Load())

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(formattedContent))
}
