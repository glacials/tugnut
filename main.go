package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/glacials/tugnut/parsers/livesplit"
	"github.com/glacials/tugnut/responses"
	"golang.org/x/net/context"
)

const (
	// memPerFile is the max amount of memory one uploaded file can take up, in bytes. If a file doesn't fit in this slot,
	// its remainder is stored in temporary files on disk. See: https://godoc.org/net/http#Request.ParseMultipartForm
	memPerFile = 1024 * 1024
)

func main() {
	ctx := context.Background()

	mux := buildMux(ctx)
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("can't start server: %s", err))
	}
}

func buildMux(ctx context.Context) http.Handler {
	mux := http.NewServeMux()

	p := livesplit.NewParser(ctx, livesplit.Config{
		ParseSegments:       true,
		ParseSegmentHistory: false,
		ParseRunHistory:     false,
	})

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
		w.Write(responses.JSONErr("Endpoints are GET /health and POST /parse/livesplit.", nil))
	}))

	mux.Handle("/health", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
	}))

	mux.Handle("/parse/livesplit", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.ParseMultipartForm(memPerFile); req.MultipartForm == nil {
			w.WriteHeader(400)
			res := responses.JSONErr(
				"You need a `splits` parameter. Make sure it's a file, not a string. In cURL: `-F splits=@/path/to/file`",
				nil,
			)
			w.Write(res)
			return
		}

		multipartStrings := req.MultipartForm.Value["splits"]
		multipartFiles := req.MultipartForm.File["splits"]

		var (
			splits io.Reader
			err    error
		)
		if len(multipartStrings) > 0 {
			splits = strings.NewReader(multipartStrings[0])
		} else if len(multipartFiles) > 0 {
			splits, err = multipartFiles[0].Open()
			if err != nil {
				w.WriteHeader(400)
				w.Write(responses.JSONErr("Received your splits file, but couldn't open it.", err))
				return
			}
		} else {
			w.Write(responses.JSONErr("You submitted a multipart form, but there was no `splits` parameter.", nil))
			return
		}

		w.WriteHeader(200)
		r, err := p.Parse(ctx, splits)
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
	return mux
}
