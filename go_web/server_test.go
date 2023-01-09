package go_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Hello World")
		fmt.Fprintln(res, req.Method)
		fmt.Fprintln(res, req.RequestURI)
	})
	mux.HandleFunc("/about", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "About Heri Hakim")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
