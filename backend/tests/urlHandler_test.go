package tests

import (
	"book-manager/handlers"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUrlHandler(t *testing.T) {
    tests := []struct {
        name           string
        requestPayload handlers.URLRequest
        expectedCode   int
        expectedBody   handlers.URLResponse
    }{
        {
            name: "Test All Operation",
            requestPayload: handlers.URLRequest{
                URL:       "https://BYFOOD.com/food-EXPeriences?query=abc/",
                Operation: "all",
            },
            expectedCode: http.StatusOK,
            expectedBody: handlers.URLResponse{
                ProcessedURL: "https://www.byfood.com/food-experiences",
            },
        },
        {
            name: "Test Canonical Operation",
            requestPayload: handlers.URLRequest{
                URL:       "https://BYFOOD.com/food-EXPeriences?query=abc/",
                Operation: "canonical",
            },
            expectedCode: http.StatusOK,
            expectedBody: handlers.URLResponse{
                ProcessedURL: "https://BYFOOD.com/food-EXPeriences",
            },
        },
		{
			name: "Test Redirection Operation",
			requestPayload: handlers.URLRequest{
				URL:       "https://EXAMPLE.com/REDIRECT",
				Operation: "redirection",
			},
			expectedCode: http.StatusOK,
			expectedBody: handlers.URLResponse{
				ProcessedURL: "https://www.byfood.com/redirect",
			},
		},
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            body, err := json.Marshal(tc.requestPayload)
            if err != nil {
                t.Fatalf("Failed to marshal request: %v", err)
            }
            req, err := http.NewRequest("POST", "/url", bytes.NewBuffer(body))
            if err != nil {
                t.Fatal(err)
            }
            rr := httptest.NewRecorder()
            handler := http.HandlerFunc(handlers.UrlHandler)

            handler.ServeHTTP(rr, req)
       
            if status := rr.Code; status != tc.expectedCode {
                t.Errorf("handler returned wrong status code: got %v want %v", status, tc.expectedCode)
            }

            var response handlers.URLResponse
            if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
                t.Fatalf("Failed to decode response: %v", err)
            }

            if response != tc.expectedBody {
                t.Errorf("handler returned unexpected body: got %v want %v", response, tc.expectedBody)
            }
        })
    }
}