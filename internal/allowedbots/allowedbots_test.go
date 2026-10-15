package allowedbots

import "testing"

func TestAllowsBot(t *testing.T) {
	tests := []struct {
		name          string
		renovateID    string
		additionalIDs string
		userType      string
		userID        string
		want          bool
	}{
		{name: "renovate bot allowed by default", userType: "Bot", userID: "29139614", want: true},
		{name: "giantswarm-release bot allowed by default", userType: "Bot", userID: "342014845", want: true},
		{name: "type match is case-insensitive", userType: "bot", userID: "342014845", want: true},
		{name: "github-actions[bot] rejected", userType: "Bot", userID: "41898282", want: false},
		{name: "unknown bot rejected", userType: "Bot", userID: "1", want: false},
		{name: "empty ID rejected", userType: "Bot", userID: "", want: false},
		{name: "allowed ID with User type rejected", userType: "User", userID: "342014845", want: false},
		{name: "allowed ID with Organization type rejected", userType: "Organization", userID: "29139614", want: false},
		{name: "legacy RENOVATE_BOT_USER_ID override still honoured", renovateID: "555", userType: "Bot", userID: "555", want: true},
		{name: "override replaces default renovate ID", renovateID: "555", userType: "Bot", userID: "29139614", want: false},
		{name: "override keeps release app", renovateID: "555", userType: "Bot", userID: "342014845", want: true},
		{name: "custom list with spaces and empty entries", additionalIDs: " 7 , ,8,", userType: "Bot", userID: "8", want: true},
		{name: "custom list replaces release app default", additionalIDs: "7", userType: "Bot", userID: "342014845", want: false},
		{name: "custom list keeps renovate", additionalIDs: "7", userType: "Bot", userID: "29139614", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.renovateID, tt.additionalIDs).AllowsBot(tt.userType, tt.userID)
			if got != tt.want {
				t.Fatalf("AllowsBot(%q, %q) = %v, want %v", tt.userType, tt.userID, got, tt.want)
			}
		})
	}
}
