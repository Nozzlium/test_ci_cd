package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /halo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Halo")
	})
	server := http.Server{
		Addr:    ":4343",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
