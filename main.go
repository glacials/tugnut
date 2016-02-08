package main

import (
	"fmt"
	"net/http"

	"github.com/glacials/tugnut/parser"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		req.ParseMultipartForm(1024 * 1024)

		multipartForm := req.MultipartForm
		if multipartForm == nil {
			w.WriteHeader(400)
			w.Write(jsonErr("You need to include a `file` parameter. Make sure it's a file, not a string.", nil))
			return
		}

		fileHeaders := multipartForm.File["file"]
		file, err := fileHeaders[0].Open()
		if err != nil {
			w.WriteHeader(400)
			w.Write(jsonErr("Couldn't read your file.", nil))
			return
		}

		p := parser.New(file)

		w.WriteHeader(200)
		p.Parse()
		w.Write(jsonRun(p))
	}))

	server := http.Server{
		Addr:    ":8000",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("can't start server: %s", err))
	}
}
