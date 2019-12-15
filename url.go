// Package net contains helper function for handling
// e.g. ip addresses or domain names
package net

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// URI struct representation
// definition: URI = scheme:[//authority]path[?query][#fragment]
type URI struct {
	Scheme    string
	Authority *Authority
	Path      string
	Query     string
	Fragment  string
}

// Authority struct representation
// definition: authority = [userinfo@]host[:port]
type Authority struct {
	UserInfo string
	Host     string
	Port     uint16
}

// IsURL returns true if string represents a valid URL
func IsURL(u string) bool {
	if IsIPAddr(u) || IsNetwork(u) || IsDomain(u) || IsFqdn(u) {
		return false
	}

	if _, err := ParseURL(u); err == nil {
		return true
	}
	return false
}

// ParseURL returns URL and empty error,
// nil and error otherwise
func ParseURL(u string) (*URI, error) {
	var f []string
	// URL = scheme:[//authority]path[?query][#fragment]
	uri := &URI{}
	//authority = [userinfo@]host[:port]
	uri.Authority = &Authority{}

	// fragment
	f = strings.SplitN(u, "#", 2)
	u = f[0]
	if len(f) == 2 {
		uri.Fragment = f[1]
	}

	// query
	f = strings.SplitN(u, "?", 2)
	u = f[0]
	if len(f) == 2 {
		uri.Fragment = f[1]
	}

	// scheme
	// case 1 (easy one): <scheme>:// exists
	f = strings.SplitN(u, "://", 2)
	if len(f) == 2 {
		uri.Scheme = f[0]
		u = f[1]
	}

	// as a side effect of removing existing :// the only slashes ("/") left
	// should all be part of the path. So we go for the path first:
	f = strings.SplitN(u, "/", 2)
	if len(f) == 2 {
		uri.Path = fmt.Sprintf("/%s", f[1])
	}
	u = f[0]

	// port
	f = strings.Split(u, ":")
	if len(f) >= 2 {
		i, err := strconv.ParseUint(f[len(f)-1], 10, 16)
		if err == nil {
			uri.Authority.Port = uint16(i)
			u = u[:strings.LastIndex(u, ":")]
		}
	}

	// NOT an IPv6-address?
	if !strings.HasPrefix(u, "[") {
		// case 2: <scheme>:// does NOT exist, we are looking for <scheme>: , e.g. mailto:
		f = strings.Split(u, ":")
		if len(f) == 2 {
			uri.Scheme = f[0]
			u = f[1]
		}
	}

	// userinfo?
	f = strings.SplitN(u, "@", 2)
	if len(f) == 2 {
		uri.Authority.UserInfo = f[1]
		u = f[0]
	}

	// Host
	// handle IPv6
	u = strings.TrimRight(strings.TrimLeft(u, "["), "]")

	if IsIPAddr(u) || IsDomain(u) || IsFqdn(u) {
		uri.Authority.Host = u
	} else {
		return uri, errors.New("error parsing host")
	}

	// no scheme found yet? fall back to http
	if uri.Scheme == "" {
		uri.Scheme = "http"
	}

	return uri, nil
}

// HostFromURL extraxts hostname from given URL
func HostFromURL(url string) (string, error) {
	u, err := ParseURL(url)
	if err != nil {
		return "", err
	}

	return u.Authority.Host, nil
}
