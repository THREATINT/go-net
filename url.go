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

	if u, err = NormaliseURLSchema(u); err == nil {

		if _, err := url.Parse(u); err == nil {

			if h, err := HostFromURL(u); err == nil {

				if IsIPAddr(h) || IsDomain(h) || IsFQDN(h) {

					return true
				}
			}
		}
	}

	return false
}

// HostFromURL extraxts hostname from given URL
func HostFromURL(u string) (string, error) {

	var err error
	var host string
	var a *url.URL

	if u, err = NormaliseURLSchema(u); err != nil {

		return "", err
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

func NormaliseURLSchema(u string) (string, error) {

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

// NormaliseURLToUnicode returns normalised URL string.
func URLToUnicode(u string) (string, error) {

	var err error
	var host string
	var unicodehost string

	if host, err = HostFromURL(u); err != nil {

		return "", err

	}

	if unicodehost, err = idna.ToUnicode(host); err != nil {
		return "", err
	}
	u = strings.Replace(u, host, unicodehost, 1)
	return u, nil
}

// URLToPunycode returns URL string in punycode
func URLToPunycode(u string) (string, error) {

	var err error
	var host string
	var unicodehost string

	if host, err = HostFromURL(u); err != nil {

		return "", err

	}

	if unicodehost, err = idna.ToASCII(host); err != nil {
		return "", err
	}
	u = strings.Replace(u, host, unicodehost, 1)
	return u, nil
}

const normaliseFlags purell.NormalizationFlags = purell.FlagRemoveDefaultPort |
	purell.FlagDecodeDWORDHost | purell.FlagDecodeHexHost | purell.FlagDecodeOctalHost |
	purell.FlagRemoveUnnecessaryHostDots | purell.FlagRemoveDuplicateSlashes |
	purell.FlagUppercaseEscapes | purell.FlagDecodeUnnecessaryEscapes | purell.FlagEncodeNecessaryEscapes | purell.FlagRemoveEmptyPortSeparator | purell.FlagSortQuery

func NormaliseURL(u string) (string, error) {

	var a *url.URL
	var err error

	if !IsURL(u) {

		return "", errors.New("not a url")
	}

	if a, err = url.Parse(u); err != nil {

		return "", err
	}

	return purell.NormalizeURL(a, normaliseFlags), nil
}
