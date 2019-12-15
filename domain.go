// Package net contains helper function for handling
// e.g. ip addresses or domain names
package net

import "strings"

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
