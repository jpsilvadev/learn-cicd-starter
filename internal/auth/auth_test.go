package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	cases := []struct {
		input         http.Header
		expectedKey   string
		expectedError bool
	}{
		{
			input:         http.Header{"Authorization": []string{"ApiKey 1234"}},
			expectedKey:   "1234",
			expectedError: false,
		},
		{
			input:         http.Header{"Authorization": []string{"ApiKey test123"}},
			expectedKey:   "test123",
			expectedError: false,
		},
		{
			input:         http.Header{"Authorization": []string{"Bearer Token"}},
			expectedKey:   "",
			expectedError: true,
		},
		{
			input:         http.Header{},
			expectedKey:   "",
			expectedError: true,
		},
	}

	for _, c := range cases {
		result, err := GetAPIKey(c.input)
		if c.expectedError {
			if err == nil {
				t.Errorf("GetAPIKey(%v) expected an error but got none", c.input)
			}	
		} else {
			if err != nil {
				t.Errorf("GetAPIKey(%v) unexpected error: %v", c.input, err)
			}
			if result != c.expectedKey {
				t.Errorf("GetAPIKey(%v) = %v; want %v", c.input, result, c.expectedKey)asdad //temp breakage
		}
	}
}
