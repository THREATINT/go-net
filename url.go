// Package net contains helper function for handling
// e.g. ip addresses or domain names
package net

import (
	"strings"
)

// IsURL returns true if string represents a valid URL
func IsURL(u string) bool {
	u = strings.TrimSpace(u)
	if IsIPAddr(u) || IsNetwork(u) || IsDomain(u) || IsFqdn(u) {
		return false
	}

	url, err := Parse(u)
	if err == nil {
		_, _, err := SplitHostPort(url)
		if err == nil {
			return true
		}
	}

	return false
}

// HostFromURL extraxts hostname from given URL
func HostFromURL(u string) (string, error) {
	url, err := Parse(u)
	if err != nil {
		return "", err
	}

	host, _, err := SplitHostPort(url)
	if err != nil {
		return "", err
	}

	return host, nil
}
