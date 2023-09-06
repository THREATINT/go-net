# go-net

## Introduction
This library contains basic validators and helpers for handling IP addresses, IP networks, 
IP ranges (from - to IP adress), domain names, FQDN (fully qualified domain names), 
and URL (unified resource locators). 

All methods that work with IP addresses are IPv4 and IPv6 compliant.

Please see the unit tests (xx_test.go) for examples on how to use this library.

## Hints
We have received feedback from software developers who were confused about how this library handles 
specific cases. Examples include but are not limited to:
* URLs:
  * We consider URLs that do not have a scheme as valid. 
  This is a clear violation of [RFC#3986](https://www.rfc-editor.org/rfc/rfc3986.txt) 
  that says *The scheme and path components are required, (...)*. 
  We are doing this because of several of our own use cases where we had to process large lists of URLs that did 
  not contain a scheme in each end every case. We neither wanted to add a scheme like `http://` by default nor simply 
  reject entries. If you, however, need a URL that fully complies to RFC#3986 
  (e.g. because you want to use it with other libraries), please call ```NormaliseURLSchema()```. 
* Domains
  * To get an idea, what a valid domain name looks like, we use the list of TLDs (top-level-domains) from 
  [IANA](https://data.iana.org/TLD/tlds-alpha-by-domain.txt). 
* FQDNs
  * _www.site.tld_ **is not** a URL but a FQDN (fully qualified domain name), because it neither has a path component
  nor a schema (`http://`, `https://`, etc.). 
  * _www.site.tld/about_ **is** a valid URL.
* IP addresses and networks
  * 10.0.0.1/32 although presented as a network address is in fact a single IPv4 address (10.0.0.1).
  * 10.0.0.1/24 is **not** an IPv4 address, but a URL! This is because we do not need a schema
    (see *URLs* above), but when the netmask is /24 the last byte must be 0. 
  * 10.0.0.0/24 is a valid network address.

If you find any other behaviour that seems odd, please double check with the code of the unit tests. 
If something still does not make sense, let us know by starting a discussion or by opening an issue here.

## License
Release under the MIT License. (see LICENSE)

## QA
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/01c46c2a6f10458f8e7f09fff5ae1915)](https://app.codacy.com/gh/THREATINT/go-net/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)