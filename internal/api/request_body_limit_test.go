package api

import "testing"

func TestRequestBodyLimit(t *testing.T) {
	t.Parallel()

	for _, path := range []string{
		"/v1/contacts/import/preview",
		"/v1/contacts/import/commit",
	} {
		if got := requestBodyLimit(path); got != contactImportMaxRequestBytes {
			t.Fatalf("requestBodyLimit(%q) = %d, want %d", path, got, contactImportMaxRequestBytes)
		}
	}

	if got := requestBodyLimit("/v1/contacts"); got != defaultMaxRequestBytes {
		t.Fatalf("ordinary request limit = %d, want %d", got, defaultMaxRequestBytes)
	}
	if contactImportMaxRequestBytes <= 50<<20 {
		t.Fatal("multipart request limit must include framing headroom above a 50 MiB file")
	}
}
