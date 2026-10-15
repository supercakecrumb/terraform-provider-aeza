package client

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestLogsOmitSensitivePayloads(t *testing.T) {
	var out bytes.Buffer
	old := log.Writer()
	log.SetOutput(&out)
	defer log.SetOutput(old)
	c := &Client{}
	secret := "synthetic-sensitive-marker"
	headers := http.Header{"Set-Cookie": []string{secret}}
	c.logRequest("POST", "https://example.invalid/?token="+secret, []byte(secret), headers)
	c.logResponse("POST", "https://example.invalid/?token="+secret, "200 OK", []byte(secret), headers, time.Second)
	c.logError("POST", "https://example.invalid/?token="+secret, fmt.Errorf("%s", secret), time.Second)
	if strings.Contains(out.String(), secret) {
		t.Fatal("sensitive value logged")
	}
	if !strings.Contains(out.String(), "200 OK") {
		t.Fatal("status diagnostics missing")
	}
}
func TestAPIErrorOmitsResponseBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		fmt.Fprint(w, "synthetic-sensitive-marker")
	}))
	defer server.Close()
	c, _ := NewClient(server.URL, "synthetic-test-key")
	err := c.NewRequest("GET", "/", nil).Do(context.Background(), nil)
	if err == nil || strings.Contains(err.Error(), "synthetic-sensitive-marker") {
		t.Fatal("response body leaked through error")
	}
}
