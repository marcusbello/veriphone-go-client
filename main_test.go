package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPhoneValidity(t *testing.T) {
	var result verifyResponse
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(&result)
	}))

	want := result.PhoneValid
	repo := newVerifyPhoneRepo(s.URL)
	got, err := repo.validatePhone(context.Background(), "15123577723", "US")

	if err != nil {
		t.Fatal(err)
	}
	if got.PhoneValid != want {
		t.Errorf("Unexpected fact returned. Got %v, want %v", got.PhoneValid, want)
	}
}

func TestVerifyHTTPFunction(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			VerifyHTTPFunction(tt.args.w, tt.args.r)
		})
	}
}

func TestHomeHTTPFunction(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HomeHTTPFunction(tt.args.w, tt.args.r)
		})
	}
}
