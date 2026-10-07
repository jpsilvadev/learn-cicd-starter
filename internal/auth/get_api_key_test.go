package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr error
	}{
		{
			name:    "authorization header missing",
			headers: http.Header{},
			wantErr: ErrNoAuthHeaderIncluded,
		},
		{
			name:    "authorization header malformed",
			headers: http.Header{"Authorization": []string{"Bearer abc"}},
			wantErr: errors.New("malformed authorization header"),
		},
		{
			name:    "api key with slashes",
			headers: http.Header{"Authorization": []string{"ApiKey a/b/c"}},
			want:    "a/b/c",
		},
		{
			name:    "api key without slashes",
			headers: http.Header{"Authorization": []string{"ApiKey abc"}},
			want:    "abc",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetAPIKey(tc.headers)
			if tc.wantErr != nil {
				if err == nil || err.Error() != tc.wantErr.Error() {
					t.Fatalf("expected error %q, got %v", tc.wantErr, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Errorf("expected API key %q, got %q", tc.want, got)
			}
		})
	}
}
