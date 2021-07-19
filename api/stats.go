package api

import (
	"fmt"
	"net/http"
)

func Stats(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Stats</h1>")
}
