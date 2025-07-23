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
		{"microsoft.abcde", false},
		{"1.2.3.4", false},
		{"президент.рф", true},                   // kremlin.ru (unicode)
		{"xn--d1abbgf6aiiy.xn--p1ai", true},      // kremlin.ru (punycode)
		{"www.президент.рф", false},              // www.kremlin.ru (unicode) -> fqdn
		{"www.xn--d1abbgf6aiiy.xn--p1ai", false}, // www.kremlin.ru (punycode) -> faqn

	}

	for _, e := range testDomains {
		r := IsDomain(e.domain)
		if r != e.expected {
			t.Errorf("%s", e.domain)
		}
	}
}
