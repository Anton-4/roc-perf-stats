package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rvcas/roc-perf-stats/server"
)

func main() {
	godotenv.Load()

	r := mux.NewRouter()

	r.HandleFunc("/benchmarks", server.NewBenchmark).Methods(http.MethodPost)
	r.HandleFunc("/benchmarks", server.GetBenchmarks).Methods(http.MethodGet)

	r.Use(server.LoggingMiddleware)

	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Listening at http://localhost:8080")

	log.Fatal(srv.ListenAndServe())
}
