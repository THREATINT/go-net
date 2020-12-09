// Package net contains helper function for handling
// e.g. ip addresses or domain names
package net

import (
	"errors"
	"strings"

	"golang.org/x/net/idna"
)

// IsDomain (domainname string) returns true is domainname is a valid domain, otherwise false
func IsDomain(domainname string) bool {
	if !IsIPAddr(domainname) {
		if !strings.Contains(domainname, "/") && !strings.Contains(domainname, ":") && !strings.Contains(domainname, " ") {
			p := strings.SplitN(domainname, ".", 2)
			if len(p) == 2 {
				for _, s := range publicSuffix {
					if p[1] == s {
						return true
					}
				}
			}
		}
	}

	return false
}

// NormaliseDomainToUnicode returns normalised domain name as Unicode
func NormaliseDomainToUnicode(domainname string) (string, error) {
	domainname = strings.TrimSpace(domainname)

	if !IsDomain(domainname) {
		return "", errors.New("invalid domain name")
	}

	domainname, err := idna.ToUnicode(domainname)
	if err != nil {
		return "", err
	}

	return domainname, nil

}

// NormaliseDomainToPunycode returns normalised domain name as Punycode
func NormaliseDomainToPunycode(domainname string) (string, error) {
	domainname = strings.TrimSpace(domainname)

	if !IsDomain(domainname) {
		return "", errors.New("invalid domain name")
	}

	domainname, err := idna.ToASCII(domainname)
	if err != nil {
		return "", err
	}

	return domainname, nil
}
