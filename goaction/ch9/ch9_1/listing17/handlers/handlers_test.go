package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/wxb/goLab/goaction/ch9/ch9_1/listing17/handlers"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

func init() {
	handlers.Routes()
}

func TestSendJSON(t *testing.T) {
	t.Log("Given the need to test the SendJSON endpoint.")

	// req := httptest.NewRequest("GET", "/sendjson", nil)
	req, err := http.NewRequest("GET", "/sendjson", nil)
	if err != nil {
		t.Fatal("\tShould be able to create a request.", ballotX, err)
	}
	t.Log("\tShould be able to create a reqeust.", checkMark)

	w := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatal("\tShould receive \"200\" ", ballotX, w.Code)
	}
	t.Log("\tShould receive \"200\" ", checkMark)

	u := struct {
		Name  string
		Email string
	}{}

	err = json.NewDecoder(w.Body).Decode(&u)
	if err != nil {
		t.Fatal("\tShould decode the response. ", ballotX)
	}
	t.Log("\tShould decode the response.", checkMark)

	if u.Name == "Bill" {
		t.Log("\tShould have a Name.", checkMark)
	} else {
		t.Error("\tShould have a Name.", ballotX, u.Name)
	}

	if u.Email == "bill@ardanstudios.com" {
		t.Log("\tShould have an Email.", checkMark)
	} else {
		t.Error("\tShould have an Email.", ballotX, u.Email)
	}
}
