package server

import (
	"io"
	"strings"
	"testing"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/humatest"
)

func TestHelloWorldHandler(t *testing.T) {
	_, api := humatest.New(t)
	s := &Server{}

	huma.Get(api, "/", s.HelloWorldHandler)

	resp := api.Get("/")

	expected := "{\"message\":\"Hello World\"}"
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading response body. Err: %v", err)
	}
	if expected != strings.TrimSpace(string(body)) {
		t.Errorf("expected response body to be %v; got %v", expected, string(body))
	}
}
