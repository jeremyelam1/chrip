package main

// type apiConfig struct {
// 	fileserverHits atomic.Int32
// }
//
// func main() {
// 	const filepathRoot = "."
// 	const port = "8080"
//
// 	apiCfg := apiConfig{
// 		fileserverHits: atomic.Int32{},
// 	}
//
// 	mux := http.NewServeMux()
// 	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))))
// 	// mux.HandleFunc("/healthz", handlerReadiness)
// 	mux.HandleFunc("/metrics", apiCfg.handlerMetrics)
// 	mux.HandleFunc("/reset", apiCfg.handlerReset)
//
// 	srv := &http.Server{
// 		Addr:    ":" + port,
// 		Handler: mux,
// 	}
//
// 	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
// 	log.Fatal(srv.ListenAndServe())
// }
//
// func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(fmt.Sprintf("Hits: %d", cfg.fileserverHits.Load())))
// }
//
// func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		cfg.fileserverHits.Add(1)
// 		next.ServeHTTP(w, r)
// 	})
// }

// func main() {
// 	mux := http.NewServeMux()
//
// 	apiConfig := NewAPIConfig()
//
// 	mux.Handle(
// 		"/app/",
// 		apiConfig.middlewareMetricsInc(
// 			http.StripPrefix("/app/", apiConfig.middlewareMetricsInc(http.FileServer(http.Dir(".")))),
// 		),
// 	)
// 	mux.Handle(
// 		"/assets",
// 		apiConfig.middlewareMetricsInc(http.StripPrefix("/app/", http.FileServer(http.Dir("./assets/logo.png")))),
// 	)
//
// 	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte("OK"))
// 	})
//
// 	mux.HandleFunc("/metrics", apiConfig.handlerMetrics)
//
// 	mux.HandleFunc("/reset", apiConfig.handlerReset)
//
// 	server := &http.Server{
// 		Addr:    ":8080",
// 		Handler: mux,
// 	}
//
// 	fmt.Println("Server started on :8080")
// 	server.ListenAndServe()
// }
//
// type apiConfig struct {
// 	fileserverHits atomic.Int32
// }
//
// func NewAPIConfig() *apiConfig {
// 	return &apiConfig{
// 		fileserverHits: atomic.Int32{},
// 	}
// }
//
// func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(fmt.Sprintf("Hits: %d", cfg.fileserverHits.Load())))
// }
//
// func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
// 	fmt.Println("testing")
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		cfg.fileserverHits.Add(1)
// 		next.ServeHTTP(w, r)
// 	})
// }
//
// func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
// 	cfg.fileserverHits.Store(0)
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Hits reset to 0"))
// }
