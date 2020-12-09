// Package net contains helper function for handling
// e.g. ip addresses or domain names
package net

import "testing"

func TestIsURL(t *testing.T) {
	var urlTests = []struct {
		url      string
		expected bool
	}{
		{"http://localhost", true},
		{"http://server", true},
		{"http://www.microsoft.com", true},
		{"microsoft.com", false},
		{"http://www.microsoft.com", true},
		{"http://microsoft.com", true},
		{"http://microsoft.com:80", true},
		{"https://microsoft.com:443/test", true},
		{"https://microsoft.com?hello#fragment", true},
		{"http://[2001:db8::1]/32", true},
		{"https://[2001:db8::1]:80/32", true},
		{"[2001:db8::1]:80/32", true},
		{"[2001:db8::]/32", true},
		{"1.2.3.4/24", true},
		{"1.2.3.0/24", false},
		{"2001:db8::/32", false},
		{"🐶🚀.ws/wp-content/5t6c5-6mz2wz6-qibv", true},
		{"xn--bp8hu0b.ws/wp-content/5t6c5-6mz2wz6-qibv", true},
	}

	for _, e := range urlTests {
		if IsURL(e.url) != e.expected {
			t.Errorf("%s", e.url)
		}
	}
}

func TestHostFromURL(t *testing.T) {
	var hostTests = []struct {
		url      string
		expected string
	}{
		{"http://localhost", "localhost"},
		{"http://www.microsoft.com", "www.microsoft.com"},
		{"microsoft.com", "microsoft.com"},
		{"http://microsoft.com", "microsoft.com"},
		{"http://microsoft.com:80", "microsoft.com"},
		{"https://www.microsoft.com:443/test", "www.microsoft.com"},
		{"https://microsoft.com?hello#fragment", "microsoft.com"},
		{"[2001:db8::1]/32", "[2001:db8::1]"},
		{"[2001:db8::1]:80/32", "[2001:db8::1]"},
		{"[2001:db8::]/32", "[2001:db8::]"},
		{"1.2.3.4/24", "1.2.3.4"},
		{"1.2.3.4:8080/24", "1.2.3.4"},
		{"https://1.2.3.4:8080/24", "1.2.3.4"},
	}

	for _, e := range hostTests {
		r, err := HostFromURL(e.url)
		if r != e.expected || err != nil {
			t.Errorf("expected: '%s' != '%s'", e.expected, r)
		}
	}
}
