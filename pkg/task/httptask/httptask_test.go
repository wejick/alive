package httptask

import (
	"net/http"
	"reflect"
	"testing"
)

func Test_ConstructURL(t *testing.T) {
	tests := []struct {
		protocol string
		domain   string
		endpoint string
		expected string
	}{
		{"http", "example.com", "/home", "http://example.com/home"},
		{"https", "example.com", "/about", "https://example.com/about"},
		{"ftp", "fileserver.com", "/downloads", "ftp://fileserver.com/downloads"},
	}

	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			result := constructURL(test.protocol, test.domain, test.endpoint)
			if result != test.expected {
				t.Errorf("Expected %s, but got %s", test.expected, result)
			}
		})
	}
}

func TestConstructHeader(t *testing.T) {
	tests := []struct {
		headersString map[string]string
		expected      http.Header
	}{
		{
			map[string]string{"Content-Type": "application/json", "Accept": "application/json"},
			http.Header{"Content-Type": []string{"application/json"}, "Accept": []string{"application/json"}},
		},
		{
			map[string]string{"User-Agent": "Go-http-client/1.1"},
			http.Header{"User-Agent": []string{"Go-http-client/1.1"}},
		},
		{
			nil,
			nil,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := constructHeader(test.headersString)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
