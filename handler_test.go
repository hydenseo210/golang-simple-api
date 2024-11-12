package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"testing/fstest"
)

func TestHandlers(t *testing.T) {

	testFS := fstest.MapFS{
		"file.go": {},
	}

	tests := []struct {
		request      http.HandlerFunc
		path         string
		expectedCode int
		expectedData string
	}{
		{
			request:      hello,
			path:         "http://localhost/hello",
			expectedCode: http.StatusOK,
			expectedData: "Hello World\n",
		},
		{
			request:      health,
			path:         "http://localhost/health",
			expectedCode: http.StatusOK,
			expectedData: "",
		},
		{
			request:      directory(testFS),
			path:         "http://localhost/directory",
			expectedCode: http.StatusOK,
			expectedData: ". file.go",
		},
		{
			request:      metadata,
			path:         "http://localhost/metadata",
			expectedCode: http.StatusOK,
			expectedData: `[{"version":"1","description":"pre-interview technical test","last_commit_sha":""}]`,
		},
	}

	for _, i := range tests {

		req := httptest.NewRequest(http.MethodGet, i.path, nil)
		w := httptest.NewRecorder()
		i.request(w, req)
		res := w.Result()
		defer res.Body.Close()
		data, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("expected error to be nil, got %v", err)
		}
		if res.StatusCode != i.expectedCode {
			t.Fatalf("Received non-200 response: %d\n", res.StatusCode)
		}
		if string(data) != i.expectedData {
			t.Errorf("expected %v, got %v", i.expectedData, string(data))
		}
	}

}
