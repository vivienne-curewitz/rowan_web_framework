package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServerHandler(t *testing.T) {
	public := getFileSystem()
	handler := getHandler(public)

	tests := []struct {
		name           string
		path           string
		expectedStatus int
		expectedBody   string // partial match
	}{
		{
			name:           "Serve index.html for root",
			path:           "/",
			expectedStatus: http.StatusOK,
			expectedBody:   "<title>Vite App</title>",
		},
		{
			name:           "Serve existing asset",
			path:           "/assets/index-DDLJ1JRl.js",
			expectedStatus: http.StatusOK,
			expectedBody:   "", // just check status
		},
		{
			name:           "Serve existing image",
			path:           "/IMG_2624.jpeg",
			expectedStatus: http.StatusOK,
			expectedBody:   "", // binary content
		},
		{
			name:           "SPA Fallback for non-existent route",
			path:           "/non-existent-route",
			expectedStatus: http.StatusOK,
			expectedBody:   "<title>Vite App</title>",
		},
		{
			name:           "SPA Fallback for nested route",
			path:           "/about/me",
			expectedStatus: http.StatusOK,
			expectedBody:   "<title>Vite App</title>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.path, nil)
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %v, got %v", tt.expectedStatus, rr.Code)
			}

			if tt.expectedBody != "" && !strings.Contains(rr.Body.String(), tt.expectedBody) {
				t.Errorf("expected body to contain %q, but it didn't", tt.expectedBody)
			}
		})
	}
}
