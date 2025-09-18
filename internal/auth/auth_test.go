package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestEmptyHeader(t *testing.T) {
	req, _ := http.NewRequest("GET", "localhost:8080", nil)
	got, _ := GetAPIKey(req.Header)
	want := ""
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestValidHeader(t *testing.T) {
	req, _ := http.NewRequest("GET", "localhost:8080", nil)
	req.Header.Add("Authorization", "ApiKey ApiValue")
	got, _ := GetAPIKey(req.Header)
	want := "ApiValue"
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
