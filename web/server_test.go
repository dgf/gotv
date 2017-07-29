package web_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dgf/gotv/memory"
	"github.com/dgf/gotv/web"
)

func TestGetStatus(t *testing.T) {
	server := httptest.NewServer(web.New("test", memory.New()))
	defer server.Close()

	response, err := http.Get(server.URL + "/status")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = response.Body.Close() }()

	if response.StatusCode != 200 {
		t.Errorf("status code: %d\n", response.StatusCode)
	}

	contentType := response.Header.Get("Content-Type")
	if contentType != "application/json; charset=UTF-8" {
		t.Errorf("content type: %s\n", contentType)
	}

	status := web.Status{}
	err = json.NewDecoder(response.Body).Decode(&status)
	if err != nil {
		t.Fatal(err)
	}
	if status.Name != "test" || status.Games != 0 {
		t.Errorf("invalid status: %#v\n", status)
	}
}
