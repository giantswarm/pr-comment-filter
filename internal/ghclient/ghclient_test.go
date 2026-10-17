package ghclient

import "testing"

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		token string
	}{
		{name: "empty token gives an unauthenticated client", token: ""},
		{name: "token gives an authenticated client", token: "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := New(tt.token)
			if err != nil {
				t.Fatalf("New(%q) returned error: %v", tt.token, err)
			}
			if c == nil {
				t.Fatalf("New(%q) returned nil client", tt.token)
			}
		})
	}
}
