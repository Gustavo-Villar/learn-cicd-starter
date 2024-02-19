package auth

import (
	"net/http"
	"testing"
)

// Mock the error for no authorization header for the sake of completeness in this example

// Your GetAPIKey function here for reference. In actuality, you'd be importing it from the package where it's defined.

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name           string
		headers        http.Header
		expectedApiKey string
		expectError    bool
	}{
		{
			name:           "No Authorization Header",
			headers:        make(http.Header),
			expectedApiKey: "",
			expectError:    true,
		},
		{
			name: "Malformed Authorization Header",
			headers: http.Header{
				"Authorization": []string{"MalformedHeader"},
			},
			expectedApiKey: "",
			expectError:    true,
		},
		{
			name: "Valid API Key",
			headers: http.Header{
				"Authorization": []string{"ApiKey validApiKey123"},
			},
			expectedApiKey: "validApiKey123",
			expectError:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.expectError {
				t.Errorf("GetAPIKey() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if apiKey != tt.expectedApiKey {
				t.Errorf("GetAPIKey() = %v, expected %v", apiKey, tt.expectedApiKey)
			}
		})
	}
}
