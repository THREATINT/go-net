# go-net

## Introduction
This library contains some basic validation and helpers for handling IP addresses, IP networks, IP ranges (from - to IP adress), domain names, FQDN (fully qualified domain names), and URL (unified ressource locators). 

All methods that work with IP addresses are IPv4 and IPv6 compliant.

Please see the unit tests (xx_test.go) for examples.

## Hints
We have received feedback from software developers in the past who were confused about how this library handles specific cases. 
* URLs:
  * We consider URLs that do not have a scheme as valid, which is a clear violation of [RFC#3986](https://www.rfc-editor.org/rfc/rfc3986.txt) that defines that "The scheme and path components are required, (...)". The reason is the problem

## License
Release under the MIT License. (see LICENSE)

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/01c46c2a6f10458f8e7f09fff5ae1915)](https://www.codacy.com/gh/THREATINT/go-net/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=THREATINT/go-net&amp;utm_campaign=Badge_Grade)
