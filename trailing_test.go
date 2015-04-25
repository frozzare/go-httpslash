package httpslash

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSlash(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.Handle("/", TrailingSlash(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, trailing!")
	})))

	defer server.Close()

	resp, err := http.Get(server.URL + "/hello")

	if err != nil {
		t.Fatal(err)
	}

	if resp.Request.URL.Path != "/hello/" {
		t.Errorf("Expected path /hello/ and not '%s'", resp.Request.URL.Path)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
	}

	expected := "Hello, trailing!"
	actual, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if expected != string(actual) {
		t.Errorf("Expected the message '%s'\n", expected)
	}
}

func TestNoSlash(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.Handle("/", NoTrailingSlash(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, no slash!")
	})))

	defer server.Close()

	resp, err := http.Get(server.URL + "/hello/")

	if err != nil {
		t.Fatal(err)
	}

	if resp.Request.URL.Path != "/hello" {
		t.Errorf("Expected path /hello and not '%s'", resp.Request.URL.Path)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
	}

	expected := "Hello, no slash!"
	actual, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if expected != string(actual) {
		t.Errorf("Expected the message '%s'\n", expected)
	}
}
