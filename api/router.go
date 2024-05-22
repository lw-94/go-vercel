package api

import (
	"fmt"
	"net/http"
)

func Listen(w http.ResponseWriter, r *http.Request) {
	// serve http request
	fmt.Printf("request: %s\n", r.URL.Path)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
