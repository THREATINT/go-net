// Package net contains helper function for handling
// e.g. ip addresses or domain names
package net

import (
	"errors"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/purell"
	"golang.org/x/net/idna"
)

// IsURL returns true if string represents a valid URL
func IsURL(u string) bool {
	u = strings.TrimSpace(u)
	if IsIPAddr(u) || IsNetwork(u) || IsDomain(u) || IsFQDN(u) {
		return false
	}

	url, err := Parse(u)
	if err == nil {
		_, _, err := SplitHostPort(url)
		if err == nil {
			return true
		}
	}

	return false
}

// HostFromURL extraxts hostname from given URL
func HostFromURL(u string) (string, error) {
	url, err := Parse(u)
	if err != nil {
		return "", err
	}

	host, _, err := SplitHostPort(url)
	if err != nil {
		return "", err
	}

	host = strings.ToLower(host)

	return host, nil
}

// Parse parses raw URL string into the net/url URL struct.
// It uses the url.Parse() internally, but it slightly changes
// its behavior:
// 1. It forces the default scheme and port to http
// 2. It favors absolute paths over relative ones, thus "example.com"
//    is parsed into url.Host instead of url.Path.
// 3. It lowercases the Host (not only the Scheme).
func Parse(rawURL string) (*url.URL, error) {
	return ParseWithDefaultScheme(rawURL, "http")
}

func ParseWithDefaultScheme(rawURL string, scheme string) (*url.URL, error) {
	rawURL = defaultScheme(rawURL, scheme)

	// Use net/url.Parse() now.
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	host, _, err := SplitHostPort(u)
	if err != nil {
		return nil, err
	}
	if err := checkHost(host); err != nil {
		return nil, err
	}

	u.Host = strings.ToLower(u.Host)
	u.Scheme = strings.ToLower(u.Scheme)

	return u, nil
}

func defaultScheme(rawURL, scheme string) string {
	// Force default http scheme, so net/url.Parse() doesn't
	// put both host and path into the (relative) path.
	if strings.Index(rawURL, "//") == 0 {
		// Leading double slashes (any scheme). Force http.
		rawURL = scheme + ":" + rawURL
	}
	if !strings.Contains(rawURL, "://") {
		// Missing scheme. Force http.
		rawURL = scheme + "://" + rawURL
	}
	return rawURL
}

var (
	domainRegexp = regexp.MustCompile(`^([a-zA-Z0-9-]{1,63}\.)*([a-zA-Z0-9-]{1,63})$`)
	ipv4Regexp   = regexp.MustCompile(`^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$`)
	ipv6Regexp   = regexp.MustCompile(`^\[[a-fA-F0-9:]+\]$`)
)

func checkHost(host string) error {
	if host == "" {
		return &url.Error{Op: "host", URL: host, Err: errors.New("empty host")}
	}

	host = strings.ToLower(host)
	if domainRegexp.MatchString(host) {
		return nil
	}

	if punycode, err := idna.ToASCII(host); err != nil {
		return err
	} else if domainRegexp.MatchString(punycode) {
		return nil
	}

	// IPv4 and IPv6.
	if ipv4Regexp.MatchString(host) || ipv6Regexp.MatchString(host) {
		return nil
	}

	return &url.Error{Op: "host", URL: host, Err: errors.New("invalid host")}
}

// SplitHostPort splits network address of the form "host:port" into
// host and port. Unlike net.SplitHostPort(), it doesn't remove brackets
// from [IPv6] host and it accepts net/url.URL struct instead of a string.
func SplitHostPort(u *url.URL) (host, port string, err error) {
	if u == nil {
		return "", "", &url.Error{Op: "host", URL: host, Err: errors.New("empty url")}
	}
	host = u.Host

	// Find last colon.
	i := strings.LastIndex(host, ":")
	if i == -1 {
		// No port found.
		return host, "", nil
	}

	// Return if the last colon is inside [IPv6] brackets.
	if strings.HasPrefix(host, "[") && strings.Contains(host[i:], "]") {
		// No port found.
		return host, "", nil
	}

	if i == len(host)-1 {
		return "", "", &url.Error{Op: "port", URL: u.String(), Err: errors.New("empty port")}
	}

	port = host[i+1:]
	host = host[:i]

	if _, err := strconv.Atoi(port); err != nil {
		return "", "", &url.Error{Op: "port", URL: u.String(), Err: err}
	}

	host = strings.ToLower(host)

	return host, port, nil
}

const normaliseFlags purell.NormalizationFlags = purell.FlagRemoveDefaultPort |
	purell.FlagDecodeDWORDHost | purell.FlagDecodeHexHost | purell.FlagDecodeOctalHost |
	purell.FlagRemoveUnnecessaryHostDots | purell.FlagRemoveDuplicateSlashes |
	purell.FlagUppercaseEscapes | purell.FlagDecodeUnnecessaryEscapes | purell.FlagEncodeNecessaryEscapes | purell.FlagRemoveEmptyPortSeparator | purell.FlagSortQuery

// NormaliseURLToUnicode returns normalised URL string.
func NormaliseURLToUnicode(u *url.URL) (string, error) {
	host, port, err := SplitHostPort(u)
	if err != nil {
		return "", err
	}
	if err := checkHost(host); err != nil {
		return "", err
	}

	// Decode Punycode.
	host, err = idna.ToUnicode(host)
	if err != nil {
		return "", err
	}

	u.Host = strings.ToLower(host)
	if port != "" {
		u.Host += ":" + port
	}
	u.Scheme = strings.ToLower(u.Scheme)

	return purell.NormalizeURL(u, normaliseFlags), nil
}

// NormaliseURLToPunycode returns normalised URL string.
func NormaliseURLToPunycode(u *url.URL) (string, error) {
	host, port, err := SplitHostPort(u)
	if err != nil {
		return "", err
	}
	if err := checkHost(host); err != nil {
		return "", err
	}

	// Convert to Punycode.
	host, err = idna.ToASCII(host)
	if err != nil {
		return "", err
	}

	u.Host = strings.ToLower(host)
	if port != "" {
		u.Host += ":" + port
	}
	u.Scheme = strings.ToLower(u.Scheme)

	return purell.NormalizeURL(u, normaliseFlags), nil
}
