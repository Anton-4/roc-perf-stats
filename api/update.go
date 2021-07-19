package api

import (
	"fmt"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Update</h1>")
}
