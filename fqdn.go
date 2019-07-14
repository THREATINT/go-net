// Package net contains helper function for handling
// e.g. ip addresses or domain names
package net

import (
	"errors"
	"net/url"
	"strings"
)

// IsFqdn (fqdn string) returns true if fqdn is a FQDN (Fully Qualified Domain Name) consiting to
// hostname + domainname + tld, otherwise false
func IsFqdn(fqdn string) bool {
	if !IsIPAddr(fqdn) {
		if !strings.Contains(fqdn, "/") && !strings.Contains(fqdn, ":") && !strings.Contains(fqdn, " ") {
			if !IsDomain(fqdn) {
				if h, err := url.Parse(fqdn); err == nil {
					// the line below is NOT a mistake: if there is no path present, url.Parse
					// comes back with an empty hostname and Path containing the hostname ...
					// strange  :-|
					if h.Host == "" && h.Path != "" {
						return true
					}
				}
			}
		}
	}

	return false
}

// DomainFromFqdn returns domain name and empty error,
// undefined string and error otherwiese
func DomainFromFqdn(fqdn string) (string, error) {
	if IsFqdn(fqdn) {
		for _, s := range publicSuffix {
			if s != "" && strings.HasSuffix(fqdn, s) {
				b := strings.Split(strings.TrimSuffix(fqdn, "."+s), ".")

				// to handle microsoft.co.uk <=> .co.uk
				// we search all public suffixes to see if we have a result
				// that also matches a public suffix
				r := b[len(b)-1] + "." + s

				for _, s2 := range publicSuffix {
					if r == s2 {
						// we have found a match in the public suffixes list
						// -> continue below, do not return this result
						goto cont
					}
				}

				return r, nil
			}

		cont:
			// continue with next suffix in range
		}
	}

	return "", errors.New("not a fqdn")
}
