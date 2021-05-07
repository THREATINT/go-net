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
		{"www_de.company.com", true},
		{"www-de.company.com", true},
		{"www.de.company.com/newsticker", false},
		{"1.2.3.4", false},
		{"612050612050612050612050612050612050-dot-onk89909.wn.r.appspot.com", true},
		{"9876543456886756565656-secondary.z19.web.core.windows.net", true},
		{"президент.рф", false},                 //kremlin.ru (unicode)
		{"xn--d1abbgf6aiiy.xn--p1ai", false},    //kremlin.ru (punycode)
		{"www.президент.рф", true},              //www.kremlin.ru (unicode)
		{"www.xn--d1abbgf6aiiy.xn--p1ai", true}, //www.kremlin.ru (punycode)
	}

	for _, e := range hostTests {
		if IsFQDN(e.hostname) != e.expected {
			t.Errorf("%s", e.hostname)
		}
	}
}

func TestDomainFromFqdn(t *testing.T) {
	var domainFromFqdnTests = []struct {
		fqdn   string
		domain string
	}{
		{"www.company.com", "company.com"},
		{"a.b.core.windows.net", "windows.net"},
		{"www.windows.co.uk", "windows.co.uk"},
		{"a.b.windows.co.uk", "windows.co.uk"},
		{"a.b.com.windows.co.uk", "windows.co.uk"},
	}

	for _, e := range domainFromFqdnTests {
		if domain := DomainFromFqdn(e.fqdn); domain != "" {
			if domain != e.domain {
				t.Errorf("%s: %s ./. %s", e.fqdn, e.domain, domain)
			}
		}
	}
}
