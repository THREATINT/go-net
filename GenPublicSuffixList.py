#!/usr/bin/python3
from collections import OrderedDict

import httplib2
import re

resp, content = httplib2.Http().request("https://publicsuffix.org/list/public_suffix_list.dat")

with open('publicSuffixList_gen.go', 'w', encoding='utf-8') as f:
    f.write("package net\n\n")

    f.write("// List of know public suffixes for domain based on\n")
    f.write("// https://publicsuffix.org/list/public_suffix_list.dat\n\n")
    f.write("// WARNING:    this is generated code, do not edit!\n")
    f.write("//             Please run GenPublicSuffixList.py to update this file\n\n\n")

    f.write("var publicSuffix = []string {\n")

    for l in str(content, 'utf-8').splitlines():
        if l == '// ===END ICANN DOMAINS===':
            break

        if not (l.startswith('//') or l.startswith('!')) and l:

            l = l[len('*.'):] if l.startswith('*.') else l

            f.write('\t"')
            f.write(l)
            f.write('",\n')

    f.write('}\n')