// Package net contains helper function for handling
// e.g. ip addresses or domain names
package net

import (
	"bytes"
	"errors"
	"net"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/purell"
	"golang.org/x/net/idna"
)

// IsURL returns true if string represents a valid URL
func IsURL(u string) bool {

	var err error

	u = strings.ToLower(strings.TrimSpace(u))

	if IsIPAddr(u) || IsNetwork(u) || IsDomain(u) || IsFQDN(u) {

		return false
	}

	if u, err = normaliseURLSchema(u); err != nil {

		return false
	}

	_, err = url.Parse(u)
	return err == nil
}

// HostFromURL extraxts hostname from given URL
func HostFromURL(u string) (string, error) {

	var err error
	var host string
	var a *url.URL

	if u, err = normaliseURLSchema(u); err != nil {

		return "", err
	}

	if !IsURL(u) {

		return "", errors.New("not a url")
	}

	if a, err = url.Parse(u); err != nil {

		return "", err
	}

	// workarounds

	//known problems with net/url, see e.g. table here: https://github.com/goware/urlx
	if a.Scheme == "" && a.Host == "" {

		host = a.Path

	} else {

		if a.Host == "" && a.Path == "" {

			host = a.Scheme

		} else {

			if host, _, err = net.SplitHostPort(a.Host); err != nil {

				host = a.Host

			}

		}
	}

	host = strings.TrimLeft(strings.TrimRight(host, "]"), "[")
	i := strings.Index(host, "/")
	if i != -1 {
		host = host[:i]
	}
	return host, nil
}

func normaliseURLSchema(u string) (string, error) {

	var i int
	var regex *regexp.Regexp

	i = strings.Index(u, "://")
	if i == -1 {

		var r bytes.Buffer

		r.WriteString("http://")
		r.WriteString(u)

		return r.String(), nil
	}

	// catch e.g. www-2.example.com/hello/https://www.example.com :
	// there is no schema at the beginning, but as part of the Path!
	if !(strings.Index(u, "/") < i || strings.Index(u, "?") < i) {

		regex = regexp.MustCompile(`^[a-zA-Z]+$`)
		if regex.FindString(u[:i]) == "" {

			return "", errors.New("existing schema is invalid")
		}
	}

	return u, nil
}

const normaliseFlags purell.NormalizationFlags = purell.FlagRemoveDefaultPort |
	purell.FlagDecodeDWORDHost | purell.FlagDecodeHexHost | purell.FlagDecodeOctalHost |
	purell.FlagRemoveUnnecessaryHostDots | purell.FlagRemoveDuplicateSlashes |
	purell.FlagUppercaseEscapes | purell.FlagDecodeUnnecessaryEscapes | purell.FlagEncodeNecessaryEscapes | purell.FlagRemoveEmptyPortSeparator | purell.FlagSortQuery

// NormaliseURLToUnicode returns normalised URL string.
func NormaliseURLToUnicode(u string) (string, error) {

	if !IsURL(u) {
		return "", errors.New("not a url")
	}

	a, err := url.Parse(u)
	if err != nil {

		return "", err
	}

	host, port, err := net.SplitHostPort(a.String())
	if err != nil {

		return "", err
	}

	// Decode Punycode.
	host, err = idna.ToUnicode(host)
	if err != nil {

		return "", err
	}
	a.Host = host

	if port != "" {

		a.Host += ":" + port
	}
	a.Scheme = strings.ToLower(a.Scheme)

	return purell.NormalizeURL(a, normaliseFlags), nil
}

// NormaliseURLToPunycode returns normalised URL string.
func NormaliseURLToPunycode(u string) (string, error) {

	if !IsURL(u) {

		return "", errors.New("not a url")
	}

	a, err := url.Parse(u)
	if err != nil {

		return "", err
	}

	host, port, err := net.SplitHostPort(a.Host)
	if err != nil {

		return "", err
	}

	// Convert to Punycode.
	host, err = idna.ToASCII(host)
	if err != nil {

		return "", err
	}

	a.Host = host
	if port != "" {

		a.Host += ":" + port
	}
	a.Scheme = strings.ToLower(a.Scheme)

	return purell.NormalizeURL(a, normaliseFlags), nil
}
