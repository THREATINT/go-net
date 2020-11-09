// Package net contains helper function for handling
// e.g. ip addresses or domain names
package net

import (
	"errors"
	"fmt"
	"strings"
)

// IsFqdn (fqdn)
// returns true if fqdn is a FQDN (Fully Qualified Domain Name)
// hostname + domainname + tld, otherwise false
func IsFqdn(fqdn string) bool {
	fqdn = strings.TrimSpace(fqdn)

	if IsIPAddr(fqdn) || IsDomain(fqdn) || strings.Contains(fqdn, "/") || strings.Contains(fqdn, "@") {
		return false
	}

	domain, err := DomainFromFqdn(fqdn)
	if err == nil && domain != "" {
		i := strings.LastIndex(fqdn, domain)
		if fqdn[:i] != "" {
			return true
		}
	}

	return false
}

// DomainFromFqdn (fqdn)
// returns domain name and empty error,
// undefined string and error otherwiese
func DomainFromFqdn(fqdn string) (string, error) {
	fqdn = strings.TrimSpace(fqdn)
	if !IsIPAddr(fqdn) && !IsDomain(fqdn) {
		for _, s := range publicSuffix {
			s = fmt.Sprintf(".%s", s)
			i := strings.LastIndex(fqdn, s)
			if i != -1 {
				return fqdn[i:], nil
			}
		}
	}

	return "", errors.New("not a FQDN")
}
