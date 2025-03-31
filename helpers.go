package net

import (
	"golang.org/x/net/idna"
	"strings"
)

// ToUnicode returns string as Unicode
func ToUnicode(s string) (string, error) {
	return idna.ToUnicode(strings.ToLower(strings.TrimSpace(s)))
}

// ToPunycode returns string as Punycode
func ToPunycode(s string) (string, error) {
	return idna.ToASCII(strings.ToLower(strings.TrimSpace(s)))
}
