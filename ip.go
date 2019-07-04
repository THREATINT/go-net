package net

import (
	"errors"
	"fmt"
	"math/big"
	"net"

	"strings"
)

func IsIPAddr(ip string) bool {
	ipaddr := net.ParseIP(ip)
	if ipaddr != nil {
		if IsIPv4(ipaddr) || IsIPv6(ipaddr) {
			return true
		}
	}

	return false
}

func IsIPv4(ip net.IP) bool {
	return strings.Count(ip.String(), ":") < 2
}

func IsIPv6(ip net.IP) bool {
	return strings.Count(ip.String(), ":") >= 2
}

func ReverseIPAddr(ip string) (string, error) {
	result := ""

	if !IsIPAddr(ip) {
		return result, errors.New("invalid IP address")
	}

	ipaddr := net.ParseIP(ip)
	if IsIPv4(ipaddr) {
		ipaddr = ipaddr.To4()
	} else {
		ipaddr = ipaddr.To16()
	}
	for i := 0; i < len(ipaddr); i++ {
		result = fmt.Sprintf("%v.%s", ipaddr[i], result)
	}

	return result, nil
}

func IsNetwork(network string) bool {
	_, ipn, err := net.ParseCIDR(network)
	if err == nil {
		// attn: comparing ipn.IP.String() to the network passed to this function
		// is important to avoid entries like 1.2.3.4/3 being detected as networks
		// while they are in fact URLs!!
		if strings.Split(network, "/")[0] == ipn.IP.String() {
			return true
		}
	}

	return false
}

func IsIPRange(r string) bool {
	f := strings.Split(r, "-")
	if len(f) == 2 {
		f[0] = strings.TrimSpace(f[0])
		f[1] = strings.TrimSpace(f[1])

		if IsIPAddr(f[0]) && IsIPAddr(f[1]) {
			return true
		}
	}

	return false
}

func IntToIP(i string) net.IP {
	var ip net.IP

	ni := big.NewInt(0)
	ni.SetString(i, 10)

	b := ni.Bytes()

	if len(b) == 4 {
		// IPv4
		ip = net.IP(b)

	} else {
		// IPv6
		b2 := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

		offset := len(b2) - len(b)
		for i := range b {
			b2[i+offset] = b[i]
		}
		ip = net.IP(b2)
	}

	return ip
}
