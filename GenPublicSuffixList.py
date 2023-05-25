#!/usr/bin/python3

import httplib2

resp, content = httplib2.Http().request("https://data.iana.org/TLD/tlds-alpha-by-domain.txt")

with open('publicSuffixList_gen.go', 'w', encoding='utf-8') as f:
    f.write("package net\n\n")

    f.write("// List of know public suffixes for domains based on\n")
    f.write("// https://data.iana.org/TLD/tlds-alpha-by-domain.txt\n\n")
    f.write("// WARNING:    this is generated code, do not edit - please run GenPublicSuffixList.py to update this file\n\n\n")

    f.write("var PublicSuffix = []string {\n")

    for l in str(content, 'utf-8').splitlines():
        if not l.startswith('#') and l:

            l = l[len('*.'):] if l.startswith('*.') else l

            f.write('\t"')
            f.write(l.lower())
            f.write('",\n')

    f.write('}\n')
