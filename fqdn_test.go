// Package net contains helper function for handling
// e.g. ip addresses or domain names
package net

import "testing"

func TestIsFqdn(t *testing.T) {
	var hostTests = []struct {
		hostname string
		expected bool
	}{
		{"www.company.com", true},
		{"*.company.com", true},
		{"company.com", false},
		{"www.de.company.com", true},
		{"www.de.company.com/newsticker", false},
		{"1.2.3.4", false},
		{"612050612050612050612050612050612050-dot-onk89909.wn.r.appspot.com", true},
		{"9876543456886756565656-secondary.z19.web.core.windows.net", true},
	}

	for _, e := range hostTests {
		if IsFQDN(e.hostname) != e.expected {
			t.Errorf("%s", e.hostname)
		}
	}
}
