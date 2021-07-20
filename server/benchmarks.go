package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/rvcas/roc-perf-stats/prisma/db"
)

func GetBenchmarks(w http.ResponseWriter, r *http.Request) {
	client := db.NewClient()

	if err := client.Connect(); err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)

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
		http.Error(w, "Server Error", http.StatusInternalServerError)

		return
	}

	result, _ := json.Marshal(benchmarks)

	fmt.Fprintf(w, "%s", result)
}

func NewBenchmark(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-API-KEY") != os.Getenv("API_KEY") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)

		return
	}

	fmt.Fprintf(w, "<h1>Update</h1>")
}
