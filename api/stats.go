package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rvcas/roc-perf-stats/prisma/db"
)

func Stats(w http.ResponseWriter, r *http.Request) {
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
