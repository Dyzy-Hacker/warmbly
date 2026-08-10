package config

import "testing"

func TestMailboxOAuthBaseURLUsesConfiguredPublicURL(t *testing.T) {
	t.Setenv("API_PUBLIC_URL", "http://localhost:8080/")

	if got, want := MailboxOAuthBaseURL(), "http://localhost:8080"; got != want {
		t.Fatalf("MailboxOAuthBaseURL() = %q, want %q", got, want)
	}
}
