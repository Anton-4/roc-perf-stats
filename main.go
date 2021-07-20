package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rvcas/roc-perf-stats/prisma/db"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/update", update)
	r.HandleFunc("/stats", stats)

	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:4000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Default().Print("Listening at http://localhost:4000")

	log.Fatal(srv.ListenAndServe())

}

func stats(w http.ResponseWriter, r *http.Request) {
	client := db.NewClient()

	if err := client.Connect(); err != nil {
		fmt.Fprintf(w, "{\"message\": \"db connection error\"}")

		return
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	benchmarks, err := client.Benchmark.FindMany().Exec(ctx)

	if err != nil {
		fmt.Fprintf(w, "{\"message\": \"error getting benchmarks\"}")

		return
	}

	result, _ := json.Marshal(benchmarks)

	fmt.Fprintf(w, "%s", result)
}

func update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Update</h1>")
}
