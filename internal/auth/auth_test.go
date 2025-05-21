package auth
import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name       string
		header     http.Header
		want       string
		wantErr    bool
		errMessage string
	}{
		{
			name:       "valid API key",
			header:     http.Header{"Authorization": []string{"ApiKey abc123"}},
			want:       "abc123",
			wantErr:    false,
		},
		{
			name:       "missing header",
			header:     http.Header{},
			want:       "",
			wantErr:    true,
			errMessage: ErrNoAuthHeaderIncluded.Error(),
		},
		{
			name:       "malformed header - wrong prefix",
			header:     http.Header{"Authorization": []string{"Bearer abc123"}},
			want:       "",
			wantErr:    true,
			errMessage: "malformed authorization header",
		},
		{
			name:       "malformed header - incomplete",
			header:     http.Header{"Authorization": []string{"ApiKey"}},
			want:       "",
			wantErr:    true,
			errMessage: "malformed authorization header",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.header)
			if (err != nil) != tt.wantErr {
				t.Fatalf("expected error: %v, got: %v", tt.wantErr, err)
			}
			if err != nil && err.Error() != tt.errMessage {
				t.Fatalf("expected error message: %q, got: %q", tt.errMessage, err.Error())
			}
			if got != tt.want {
				t.Errorf("expected: %q, got: %q", tt.want, got)
			}
		})
	}
}

