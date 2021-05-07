// Package net contains helper function for handling
// e.g. ip addresses or domain names
package net

import (
	"errors"
	"fmt"
	"strings"

	"golang.org/x/net/idna"
)

// IsFQDN (fqdn) returns true if fqdn is a FQDN (Fully Qualified Domain Name) hostname + domainname + tld, otherwise false
func IsFQDN(fqdn string) bool {

	fqdn = strings.ToLower(strings.TrimSpace(fqdn))

	if IsIPAddr(fqdn) || IsDomain(fqdn) || strings.Contains(fqdn, "/") || strings.Contains(fqdn, "@") || strings.Contains(fqdn, ":") || strings.Contains(fqdn, "\\") {

		return false
	}

	if domain := DomainFromFqdn(fqdn); domain != "" {

		i := strings.LastIndex(fqdn, domain)
		if fqdn[:i] != "" {

			return true

		}
	}

	return false
}

// DomainFromFqdn returns domain name or empty string
func DomainFromFqdn(fqdn string) string {

	fqdn = strings.TrimSpace(fqdn)
	fqdn = strings.ToLower(fqdn)
	domain := ""

	if !IsIPAddr(fqdn) && !IsDomain(fqdn) {

		for _, s := range publicSuffix {

			s = fmt.Sprintf(".%s", s)
			if strings.HasSuffix(fqdn, s) {

				if i := strings.LastIndex(fqdn, s); i != -1 {

					if j := strings.LastIndex(fqdn[:i], "."); j != -1 {

						domain = fqdn[j+1:]

					}
				}
			}
		}
	}

	return domain
}

// NormaliseFQDNToUnicode returns normalised domain name as Unicode
func NormaliseFQDNToUnicode(fqdn string) (string, error) {

	fqdn = strings.TrimSpace(fqdn)
	fqdn = strings.ToLower(fqdn)

	if !IsFQDN(fqdn) {

		return "", errors.New("invalid FQDN")

	}

	fqdn, err := idna.ToUnicode(fqdn)
	if err != nil {

		return "", err

	}

	return fqdn, nil
}

// NormaliseFQDNToPunycode returns normalised domain name as Punycode
func NormaliseFQDNToPunycode(fqdn string) (string, error) {

	fqdn = strings.TrimSpace(fqdn)
	fqdn = strings.ToLower(fqdn)

	if !IsFQDN(fqdn) {

		return "", errors.New("invalid FQDN")

	}

	fqdn, err := idna.ToASCII(fqdn)
	if err != nil {

		return "", err

	}

	return fqdn, nil
}
