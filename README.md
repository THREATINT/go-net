# go-net

## Introduction
This library contains basic validators and helpers for handling IP addresses, IP networks, IP ranges (from - to IP adress), domain names, FQDN (fully qualified domain names), and URL (unified resource locators). 

All methods that work with IP addresses are IPv4 and IPv6 compliant.

Please see the unit tests (xx_test.go) for examples on how to use this library.

## Hints
We have received feedback from software developers in the past who were confused about how this library handles specific cases. Examples include but are not limited to:
* URLs:
  * We consider URLs that do not have a scheme as valid. This is a clear violation of [RFC#3986](https://www.rfc-editor.org/rfc/rfc3986.txt) that defines that "The scheme and path components are required, (...)". 
  We are doing this because of several of our own use cases where we had to process large lists of URLs that did not contain a scheme in each end every case. We neither wanted to add a scheme like http:// by default nor simply reject entries.
  If you, however, need a URL that is fully compliant to RFC#3986 (e.g. because you want to use it with other libraries), please call ```NormaliseURLSchema()```. 
* Domains
  * To get an idea, what a valid domain name looks like, we use the list from [publicsuffix.org](https://publicsuffix.org). 
* FQDNs
  * _www.website.tld_ is not a URL but a FQDN (fully qualified domain name), because it is missing the path component. 
  * _www.website.tld/about_ is a valid URL.
* IP addresses and networks
  * 10.0.0.1/32 although presented as a network address is in fact an IPv4 address (10.0.0.1).
  * 10.0.0.1/24 not not an IPv4 address, but a URL! 
  * 10.0.0.0/24 is a network address.

If you find any other behaviour that seems odd, please double check with the code of the unit tests. If something still does not make sense, let us know by starting a discussion or by opening an issue here.

## License
Release under the MIT License. (see LICENSE)

## QA
[![DeepSource](https://deepsource.io/gh/THREATINT/go-net.svg/?label=active+issues&show_trend=true&token=4HqprlCLNf-rAsnx8mVx_RMc)](https://deepsource.io/gh/THREATINT/go-net/?ref=repository-badge)

