package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello World")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	bodyString := string(body)

	fmt.Println(recorder.Code)
	fmt.Println(bodyString)
}
