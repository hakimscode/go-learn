package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}

	firstName := req.PostForm.Get("first_name")
	lastName := req.PostForm.Get("last_name")

	// ? can also use PostFormValues without parsing the body
	// ? PostFormValues already parsing the body inside it's function
	// firstName := req.PostFormValue("first_name")
	// lastName := req.PostFormValue("last_name")

	fmt.Fprintf(res, "Halo %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=Heri&last_name=Hakim")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
