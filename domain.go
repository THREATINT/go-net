package net

import (
	"strings"
)

// IsDomain (domainname string) returns true if domainname is a valid domain, otherwise false
func IsDomain(domain string) bool {
	var (
		err error
	)

	domain = strings.ToLower(strings.TrimSpace(domain))

	if strings.Contains(domain, "/") || strings.Contains(domain, "@") ||
		strings.Contains(domain, ":") || strings.Contains(domain, "\\") {
		return false
	}

	if domain, err = ToPunycode(domain); err == nil {
		for _, s := range PublicSuffix {
			if strings.HasSuffix(domain, "."+s) {
				if len(strings.Split(domain, "."))-len(strings.Split(s, ".")) == 1 {
					return true
				}
			}
		}
	}

	return false
}
