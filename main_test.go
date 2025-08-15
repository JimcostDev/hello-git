package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRootHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("¡Hola Mundo desde Go en Docker!\n"))
	}

	handler(w, req)

	resp := w.Result()
	body := w.Body.String()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("esperado status 200, obtuvo %d", resp.StatusCode)
	}
	if !strings.Contains(body, "Hola Mundo") {
		t.Errorf("esperado body conteniendo 'Hola Mundo', obtuvo %q", body)
	}
}
