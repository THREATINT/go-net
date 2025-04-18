package net

import (
	"testing"
)

func TestIsURL(t *testing.T) {
	var urlTests = []struct {
		url      string
		expected bool
	}{
		// {"http://localhost", true},
		// {"http://server", true},
		{"http://www.microsoft.com", true},
		{"HttP://www.microsoft.com", true},

		{"www_test.microsoft.com", false},

		{"http://www_test.microsoft.com", true},

		{"www-test.microsoft.com", false},

		{"microsoft.com", false},

		{"http://www-test.microsoft.com", true},

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
		{"1.2.3.0/24", false}, // -> network!!!

		{"2001:db8::/32", false},

		{"WWW.EXAMPLE.COM", false},
		{"WWW.EXAMPLE.COM/test", true},

		{"президент.рф", false},              // kremlin.ru (unicode)
		{"xn--d1abbgf6aiiy.xn--p1ai", false}, // kremlin.ru (punycode)
		{"www.президент.рф", false},
		{"www.xn--d1abbgf6aiiy.xn--p1ai", false},
		{"президент.рф/test", true},
		{"xn--d1abbgf6aiiy.xn--p1ai/test", true},
		{"www.президент.рф/test", true},
		{"www.xn--d1abbgf6aiiy.xn--p1ai/test", true},
		{"www.президент.рф:8443/test", true},
		{"www.xn--d1abbgf6aiiy.xn--p1ai:8443/test", true},
		{"https://www.xn--d1abbgf6aiiy.xn--p1ai/test", true},
		{"HTTPS://www.xn--d1abbgf6aiiy.xn--p1ai/test", true},

		{"www-2.ext.example.com:8443/hello/https://www.example.com", true},
		{"www-2.example.com/hello/https://www.example.com", true},

		{"na01.safelinks.protection.outlook.com/?url=http://enbau.net/client/past-due-invoice", true},
		{"linkprotect.cudasvc.com/url?a=http://irissnuances.com/aug2018/us/invoice-35443454&c=e", true},
		{"www.trickyguy.com/wp-includes/01-56889677218-6377383240704407401.php/https://my.klarna.com/uk/business", true},
	}

	for i, e := range urlTests {
		if IsURL(e.url) != e.expected {
			t.Errorf("%d. %s", i+1, e.url)
		}
	}
}

func TestHostFromURL(t *testing.T) {
	var hostTests = []struct {
		url      string
		expected string
	}{
		// {"http://localhost", "localhost"},
		{"http://www.microsoft.com", "www.microsoft.com"},
		{"http://microsoft.com", "microsoft.com"},
		{"http://microsoft.com:80", "microsoft.com"},
		{"https://www.microsoft.com:443/test", "www.microsoft.com"},
		{"https://microsoft.com?hello#fragment", "microsoft.com"},
		{"[2001:db8::1]/32", "2001:db8::1"},
		{"[2001:db8::1]:80/32", "2001:db8::1"},
		{"[2001:db8::]/32", "2001:db8::"},
		{"1.2.3.4/24", "1.2.3.4"},
		{"1.2.3.4:8080/24", "1.2.3.4"},
		{"https://1.2.3.4:8080/24", "1.2.3.4"},
		{"HTTPS://1.2.3.4:8080/24", "1.2.3.4"},

		{"na01.safelinks.protection.outlook.com/?url=http://enbau.net/client/past-due-invoice", "na01.safelinks.protection.outlook.com"},
		{"linkprotect.cudasvc.com/url?a=http://irissnuances.com/aug2018/us/invoice-35443454&c=e", "linkprotect.cudasvc.com"},
		{"www.trickyguy.com/wp-includes/01-56889677218-6377383240704407401.php/https://my.klarna.com/uk/business", "www.trickyguy.com"},
	}

	for _, e := range hostTests {
		r, err := HostFromURL(e.url)
		if r != e.expected || err != nil {
			t.Errorf("%s, expected: '%s' != '%s'", e.url, e.expected, r)
		}
	}
}

func TestNormaliseURLSchema(t *testing.T) {
	var hostTests = []struct {
		url      string
		expected string
	}{
		{"www.example.com/test", "http://www.example.com/test"},
		{"http://www.example.com/test", "http://www.example.com/test"},
		{"президент.рф/test", "http://президент.рф/test"},
		{"http://президент.рф/test", "http://президент.рф/test"},
		{"xn--d1abbgf6aiiy.xn--p1ai/test", "http://xn--d1abbgf6aiiy.xn--p1ai/test"},
		{"http://xn--d1abbgf6aiiy.xn--p1ai/test", "http://xn--d1abbgf6aiiy.xn--p1ai/test"},
	}

	for _, e := range hostTests {
		r, err := NormaliseURLSchema(e.url)
		if r != e.expected || err != nil {
			t.Errorf("%s, expected: '%s' != '%s'", e.url, e.expected, r)
		}
	}
}

func TestToUnicode(t *testing.T) {
	var hostTests = []struct {
		url      string
		expected string
	}{
		{"http://xn--d1abbgf6aiiy.xn--p1ai/test", "http://президент.рф/test"},
	}

	for _, e := range hostTests {
		r, err := URLToUnicode(e.url)
		if r != e.expected || err != nil {
			t.Errorf("%s, expected: '%s' != '%s'", e.url, e.expected, r)
		}
	}
}

func TestToPunycode(t *testing.T) {
	var hostTests = []struct {
		url      string
		expected string
	}{
		{"http://президент.рф/test", "http://xn--d1abbgf6aiiy.xn--p1ai/test"},
	}

	for _, e := range hostTests {
		r, err := URLToPunycode(e.url)
		if r != e.expected || err != nil {
			t.Errorf("%s, expected: '%s' != '%s'", e.url, e.expected, r)
		}
	}
}
