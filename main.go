package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

var (
	halo   string
	berapa string
)

func main() {
	if halo == "" {
		panic(
			errors.New(
				"lah mana itu halo-nya?",
			),
		)
	}

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
					"weleh",
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
