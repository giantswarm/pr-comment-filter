// Package allowedbots decides which GitHub bot accounts may trigger pipelines.
package allowedbots

import "strings"

const (
	// DefaultRenovateBotUserID is the user ID of the Renovate bot.
	DefaultRenovateBotUserID = "29139614"
	// DefaultReleaseAppBotUserID is the user ID of the giantswarm-release
	// GitHub App bot (giantswarm-release[bot]).
	DefaultReleaseAppBotUserID = "342014845"
)

// Allowlist is a set of GitHub bot user IDs.
type Allowlist map[string]struct{}

// New builds an allowlist from the Renovate bot ID (RENOVATE_BOT_USER_ID,
// kept for backwards compatibility) and a comma-separated list of additional
// bot IDs (ALLOWED_BOT_USER_IDS). Empty values fall back to the defaults:
// Renovate defaults to DefaultRenovateBotUserID and the additional list
// defaults to DefaultReleaseAppBotUserID. The two sources are combined.
func New(renovateID, additionalIDs string) Allowlist {
	if strings.TrimSpace(renovateID) == "" {
		renovateID = DefaultRenovateBotUserID
	}
	if strings.TrimSpace(additionalIDs) == "" {
		additionalIDs = DefaultReleaseAppBotUserID
	}

	a := Allowlist{}
	for _, id := range append([]string{renovateID}, strings.Split(additionalIDs, ",")...) {
		if id = strings.TrimSpace(id); id != "" {
			a[id] = struct{}{}
		}
	}
	return a
}

// AllowsBot reports whether the account is of type Bot and its ID is allowed.
// Matching is on the numeric ID only, never on the login.
func (a Allowlist) AllowsBot(userType, userID string) bool {
	if !strings.EqualFold(userType, "bot") {
		return false
	}
	_, ok := a[userID]
	return ok
}
