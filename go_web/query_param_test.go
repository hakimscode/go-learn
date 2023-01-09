package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(res, "What is your name?")
	} else {
		fmt.Fprintf(res, " Hello %s", name)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=Heri", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func SetTitle(gender string) string {
	switch gender {
	case "Male":
		return "Mr. "
	case "Female":
		return "Mrs. "
	default:
		return ""
	}
}

func SayHelloMulti(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	gender := req.URL.Query().Get("gender")

	if name == "" {
		fmt.Fprint(res, "What is your name?")
	} else {
		title := SetTitle(gender)
		fmt.Fprintf(res, " Hello %s %s", title, name)
	}
}

func TestQueryMultiParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=Heri&gender=Male", nil)
	recorder := httptest.NewRecorder()

	SayHelloMulti(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func SayHelloMultiValues(res http.ResponseWriter, req *http.Request) {
	queryParam := req.URL.Query()
	names := queryParam["name"]

	fullName := strings.Join(names, " ")

	fmt.Fprint(res, string(fullName))
}

func TestQueryMultiValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=Heri&name=Hakim&name=Setiawan", nil)
	recorder := httptest.NewRecorder()

	SayHelloMultiValues(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
