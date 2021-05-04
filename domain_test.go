// Package net contains helper function for handling
// e.g. ip addresses or domain names
package net

import "testing"

func TestIsDomain(t *testing.T) {
	var testDomains = []struct {
		domain   string
		expected bool
	}{
		{"microsoft.com", true},
		{"microsoft.co.uk", true},
		{"www.microsoft.com", false},
		{"microsoft.com.abc", false},
		{"1.2.3.4", false},
		{"президент.рф", true}, //kremlin.ru
	}

	for _, e := range testDomains {
		r := IsDomain(e.domain)
		if r != e.expected {
			t.Errorf("%s", e.domain)
		}
	}
}
