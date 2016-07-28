package main

import (
	"fmt"
	"net/http"

	"github.com/glacials/tugnut/internal"
	"github.com/glacials/tugnut/parsers/splits"
)

func main() {
	mux := buildMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("can't start server: %s", err))
	}
}

func buildMux() http.Handler {
	mux := http.NewServeMux()

	p, err := splits.NewParser(splits.LiveSplit)
	if err != nil {
		panic(fmt.Sprintf("can't make a parser: %s", err))
	}

	internal.Parser = p

	mux.Handle("/", http.HandlerFunc(internal.Root))
	mux.Handle("/health", http.HandlerFunc(internal.HealthCheck))
	mux.Handle("/parse/livesplit", http.HandlerFunc(internal.ParseLiveSplit))
	mux.Handle("/parse/livesplit/layout", http.HandlerFunc(internal.ParseLiveSplitLayout))
	return mux
}
