package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Hello_table(t *testing.T) {
	type arg struct {
		method     string
		stringBody string
	}

	type want struct {
		statusCode int
		content    string
	}

	tests := []struct {
		name string
		arg  arg
		want want
	}{
		{
			name: "wrong http method",
			arg:  arg{method: http.MethodPut, stringBody: ""},
			want: want{statusCode: http.StatusMethodNotAllowed, content: ""},
		},
		{
			name: "names is empty",
			arg:  arg{method: http.MethodGet, stringBody: ""},
			want: want{statusCode: http.StatusOK, content: "{\"names\":null}\n"},
		},
		{
			name: "name daniel is added",
			arg:  arg{method: http.MethodPost, stringBody: "{\"name\":\"daniel\"}\n"},
			want: want{statusCode: http.StatusCreated, content: "{\"message\":\"Hello, daniel!\",\"exists\":false}\n"},
		},
		{
			name: "names has daniel",
			arg:  arg{method: http.MethodGet, stringBody: ""},
			want: want{statusCode: http.StatusOK, content: "{\"names\":[\"daniel\"]}\n"},
		},
		{
			name: "name daniel already exist",
			arg:  arg{method: http.MethodPost, stringBody: "{\"name\":\"daniel\"}\n"},
			want: want{statusCode: http.StatusOK, content: "{\"message\":\"Hello, daniel! Welcome back!\",\"exists\":true}\n"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req, err := http.NewRequest(tt.arg.method, "hello", bytes.NewBuffer([]byte(tt.arg.stringBody)))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")
			res := httptest.NewRecorder()
			handler := http.HandlerFunc(helloHandler)

			handler.ServeHTTP(res, req)

			if tt.want.statusCode != res.Code {
				t.Errorf("Status Code = %d, want %d", res.Code, tt.want.statusCode)
			}

			if tt.want.content != res.Body.String() {
				t.Errorf("Content = %s, want %s", res.Body.String(), tt.want.content)
			}
		})
	}
}
