// Package ghclient builds the GitHub API client.
package ghclient

import "github.com/google/go-github/v92/github"

// New returns a GitHub client authenticated with token. An empty token yields
// an unauthenticated client, as go-github's WithAuthToken rejects "".
func New(token string) (*github.Client, error) {
	if token == "" {
		return github.NewClient()
	}
	return github.NewClient(github.WithAuthToken(token))
}
