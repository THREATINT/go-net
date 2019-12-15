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
		{"www.microsoft.com", false},
		{"microsoft.com.abc", false},
		{"1.2.3.4", false},
	}

	for _, e := range testDomains {
		r := IsDomain(e.domain)
		if r != e.expected {
			t.Errorf("%s", e.domain)
		}
	}
}

func TestDomainFromFqdn(t *testing.T) {
	var urlTests = []struct {
		url      string
		expected string
	}{
		{"www.microsoft.com", "microsoft.com"},
		{"www.microsoft.co.uk", "microsoft.co.uk"},
	}

	for _, e := range urlTests {
		r, err := DomainFromFqdn(e.url)
		if r != e.expected && err != nil {
			t.Errorf("%s != %s", e.url, e.expected)
		}
	}
}
