#!/usr/bin/python3

import httplib2

resp, content = httplib2.Http().request("https://publicsuffix.org/list/public_suffix_list.dat")

with open('publicSuffixList_gen.go', 'w', encoding='utf-8') as f:
    f.write("package net\n\n")

    f.write("// List of know public suffixes for domain based on\n")
    f.write("// https://publicsuffix.org/list/public_suffix_list.dat\n\n")
    f.write("// WARNING:    this is generated code, do not edit!\n")
    f.write("//             Please run GenPublicSuffixList.py to update this file\n\n\n")

    f.write("var PublicSuffix = []string {\n")

    for l in str(content, 'utf-8').splitlines():
        l = l.strip()

        if l == '// ===END ICANN DOMAINS===':
            break

        if l.startswith('//') or l.startswith('!') or l == '':
            continue

        l = l[len('*.'):] if l.startswith('*.') else l

        if l == 'arpa' or l.endswith('.arpa'):
            continue

        f.write('\t"')
        p = l.encode('idna')
        if l != p:
            f.write(p.decode('utf-8'))
        else:
            f.write(l)
        f.write('",\n')

    f.write('\t"bit",\n') # .bit (Bitcoin)

    f.write('}\n')