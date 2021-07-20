package server

import (
	"context"
	"encoding/json"
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

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func NewBenchmark(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-API-KEY") != os.Getenv("API_KEY") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)

		return
	}

	b := db.BenchmarkModel{}

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, "Error decoding response object", http.StatusBadRequest)

		return
	}

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

	_, err := client.Benchmark.CreateOne(
		db.Benchmark.Name.Set(b.Name),
		db.Benchmark.LowBoundConfInterval.Set(b.LowBoundConfInterval),
		db.Benchmark.LowBoundUnit.Set(b.LowBoundUnit),
		db.Benchmark.PointEstimate.Set(b.PointEstimate),
		db.Benchmark.PointEstimateUnit.Set(b.PointEstimateUnit),
		db.Benchmark.UpBoundConfInterval.Set(b.UpBoundConfInterval),
		db.Benchmark.UpBoundUnit.Set(b.UpBoundUnit),
		db.Benchmark.NrOfMeasurements.Set(b.NrOfMeasurements),
		db.Benchmark.NrOfOutliers.Set(b.NrOfOutliers),
		db.Benchmark.StartTime.Set(b.StartTime),
		db.Benchmark.Commit.Set(b.Commit),
		db.Benchmark.Branch.Set(b.Branch),
	).Exec(ctx)

	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)

		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
