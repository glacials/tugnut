package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/glacials/tugnut/parser"
	"github.com/glacials/tugnut/responses"
	"github.com/glacials/tugnut/run"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		req.ParseMultipartForm(1024 * 1024)

		multipartForm := req.MultipartForm
		if multipartForm == nil {
			w.WriteHeader(400)
			w.Write(responses.JSONErr("You need to include a `file` parameter. Make sure it's a file, not a string.", nil))
			return
		}

		fileHeaders := multipartForm.File["file"]
		file, err := fileHeaders[0].Open()
		if err != nil {
			w.WriteHeader(400)
			w.Write(responses.JSONErr("Couldn't read your file.", err))
			return
		}

		p := parser.New(run.Config{
			Parsables: map[run.Parsable]struct{}{
				run.History:        struct{}{},
				run.Segments:       struct{}{},
				run.SegmentHistory: struct{}{},
			},
		})

		w.WriteHeader(200)
		r, err := p.Parse(file)
		if err != nil {
			w.WriteHeader(400)
			w.Write(responses.JSONErr("Couldn't parse your file.", err))
		}

		j, err := json.Marshal(r)
		if err != nil {
			w.WriteHeader(400)
			w.Write(responses.JSONErr("Run was valid, but encountered an error preparing it.", err))
			return
		}

		w.Write(j)
	}))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("can't start server: %s", err))
	}
}
