package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey1(t *testing.T) {
	header := make(http.Header)
	header.Add("Authorization", "ApiKey 12345")
	got, err := GetAPIKey(header)
	if err != nil {
		t.Fatalf("function error")
	}
	want := "12345"
	if want != got {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestGetAPIKey2(t *testing.T) {
	header := make(http.Header)
	header.Add("Authorization", "ApiKey")
	_, err := GetAPIKey(header)
	if err != nil {
		if err.Error() != "malformed authorization header" {
			t.Error(err)
		}
	}
}
