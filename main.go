package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/nozzlium/test_ci_cd/lib"
)

func main() {
	lib.ReadEnv()
	halo := os.Getenv("HALO")
	if halo == "" {
		panic(
			errors.New(
				"lah mana itu halo-nya?",
			),
		)
	}

	berapa := os.Getenv("BERAPA")
	berapaNum, err := strconv.Atoi(
		berapa,
	)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc(
		"GET /halo",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(
				w,
				fmt.Sprint(
					"ini sudah berubah",
					halo,
					berapaNum,
				),
			)
		},
	)
	server := http.Server{
		Addr:    ":4343",
		Handler: mux,
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
