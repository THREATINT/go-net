package net

import (
	"fmt"
	"strings"
)

// IsFQDN (fqdn) returns true if fqdn is a FQDN (Fully Qualified Domain Name) hostname + domainname + tld,
// otherwise false
func IsFQDN(fqdn string) bool {
	var (
		err error
	)

	fqdn = strings.ToLower(strings.TrimSpace(fqdn))

	if fqdn, err = ToPunycode(fqdn); err == nil {
		if IsIPAddr(fqdn) || IsDomain(fqdn) ||
			strings.Contains(fqdn, "/") || strings.Contains(fqdn, "@") ||
			strings.Contains(fqdn, ":") || strings.Contains(fqdn, "\\") {
			return false
		}

		if domain := DomainFromFqdn(fqdn); domain != "" {
			i := strings.LastIndex(fqdn, domain)
			if fqdn[:i] != "" {
				return true
			} else {
				fmt.Printf("Und raus! %s\n", fqdn)
			}
		}
	}

	return false
}

// DomainFromFqdn returns domain name or empty string
func DomainFromFqdn(fqdn string) string {
	var (
		domain string
	)

	fqdn = strings.ToLower(strings.TrimSpace(fqdn))

	if !IsIPAddr(fqdn) && !IsDomain(fqdn) {
		for _, s := range PublicSuffix {
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
