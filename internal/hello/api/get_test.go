package hello

import (
	oapicodegen "openapi/internal/infra/oapi_codegen/hello"
	"testing"

	cmp "github.com/google/go-cmp/cmp"

	"encoding/json"
	"io"
	"net/http"
)

// TestHello is a function to test the hello API.
//
// It takes a testing.T parameter and has no return type.
func TestHello(t *testing.T) {
	res, err := http.Get("http://localhost:3000/")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, res.StatusCode)
	}

	resBodyByte, _ := io.ReadAll(res.Body)
	var actual = &oapicodegen.Hello{}
	json.Unmarshal(resBodyByte, &actual)

	var expect = &oapicodegen.Hello{
		Message: "Hello, World!",
	}

	if !cmp.Equal(actual, expect) {
		t.Errorf("expected %s, actual %s", expect, actual)
	}
}
