package net

import "testing"

func TestIsFqdn(t *testing.T) {
	var hostTests = []struct {
		hostname string
		expected bool
	}{
		{"www.heise.de", true},
		{"heise.de", false},
		{"www.heise.de/newsticker", false},
		{"1.2.3.4", false},
	}

	for _, e := range hostTests {
		if IsFqdn(e.hostname) != e.expected {
			t.Errorf("%s", e.hostname)
		}
	}
}
