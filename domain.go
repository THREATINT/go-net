// Package net contains helper function for handling
// e.g. ip addresses or domain names
package net

import (
	"errors"
	"strings"

	"golang.org/x/net/idna"
)

// IsDomain (domainname string) returns true if domainname is a valid domain, otherwise false
func IsDomain(domainname string) bool {
	domainname = strings.TrimSpace(domainname)
	domainname = strings.ToLower(domainname)

	if !IsIPAddr(domainname) {
		if !strings.Contains(domainname, "/") && !strings.Contains(domainname, ":") && !strings.Contains(domainname, " ") {
			p := strings.SplitN(domainname, ".", 2)
			if len(p) == 2 {
				for _, s := range PublicSuffix {
					if p[1] == s {
						return true
					}
				}
			}
		}
	}

	return false
}

// DomainToUnicode returns  domain name as Unicode
func DomainToUnicode(domainname string) (string, error) {
	domainname = strings.TrimSpace(domainname)
	domainname = strings.ToLower(domainname)

	if !IsDomain(domainname) {
		return "", errors.New("invalid domain name")
	}

	return idna.ToUnicode(domainname)
}

// DomainToPunycode returns normalised domain name as Punycode
func DomainToPunycode(domainname string) (string, error) {
	domainname = strings.TrimSpace(domainname)
	domainname = strings.ToLower(domainname)

	if !IsDomain(domainname) {
		return "", errors.New("invalid domain name")
	}

	return idna.ToASCII(domainname)
}
