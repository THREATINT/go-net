// Package net contains helper function for handling
// e.g. ip addresses or domain names
package net

import (
	"testing"
)

func TestIsIPv4(t *testing.T) {
	var testIPs = []struct {
		ip       string
		expected bool
	}{
		{"1.2.3.4", true},
		{"256.0.0.1", false},
		{"1.2.3.4.5", false},
		{"0.175", false},
		{".12", false},
	}

	for _, e := range testIPs {
		r := IsIPAddr(e.ip)
		if r != e.expected {
			t.Errorf("%s", e.ip)
		}
	}
}

func TestIsIPv6(t *testing.T) {

	var testIPs = []struct {
		ip       string
		expected bool
	}{
		{"1.2.3.4", true},
		{"256.0.0.1", false},
		{"1.2.3.4.5", false},
	}

	for _, e := range testIPs {
		r := IsIPAddr(e.ip)
		if r != e.expected {
			t.Errorf("%s", e.ip)
		}
	}
}

func TestIsNetwork(t *testing.T) {
	var networkTests = []struct {
		network  string
		expected bool
	}{
		{"1.2.3.0/24", true},
		{"2001:db8::/32", true},
		{"1.2.3.4", false},
		{"1.2.3.4/24", false},
		{"2001:db8::1/32", false},
	}

	for _, e := range networkTests {
		if IsNetwork(e.network) != e.expected {
			t.Errorf("%s", e.network)
		}
	}
}

func TestIsIPRange(t *testing.T) {
	var iprangeTests = []struct {
		iprange  string
		expected bool
	}{
		{"1.2.3.1-1.2.3.255", true},
		{"1.2.3.1-1.2.3.255", true},
		{"2001:db8::1-2001:db8::ffff", true},
		{"1.2.3.4", false},
		{"1.2.3.4-100", false},
	}

	for _, e := range iprangeTests {
		if IsIPRange(e.iprange) != e.expected {
			t.Errorf("%s", e.iprange)
		}
	}
}

func TestReverseIP(t *testing.T) {
	//google-public-dns-b.google.com
	r, err := ReverseIPAddr("8.8.4.4")
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%s", r)
	}

	r, err = ReverseIPAddr("2001:4860:4860::8844")
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%s", r)
	}
}

/*
func TestIsURL(t *testing.T)  { }
*/

func TestIntToIP(t *testing.T) {

	s := "16777216"
	ip := IntToIP(s)
	if ip.String() != "1.0.0.0" {
		t.Errorf("%s - %s", s, ip.String())
	}

	s = "281470698520576"
	ip = IntToIP(s)
	if ip.String() != "1.0.0.0" {
		t.Errorf("%s - %s", s, ip.String())
	}

	s = "340277174624079928635746358406137511936"
	ip = IntToIP(s)
	if ip.String() != "ffff::ffff:100:0" {
		t.Errorf("%s - %s", s, ip.String())
	}
}
