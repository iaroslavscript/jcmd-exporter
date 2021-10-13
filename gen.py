#!/usr/bin/env python3

import re


def get_name(s):
    result = []

    for c in s:
        if c.isupper():
            result.append('_')
            c = c.lower()
        
        result.append(c.lower())

    return ''.join(result)


def get_help(s):
    result = list("jcmd VM.native_memory metric")

    for c in s:
        if c.isupper():
            result.append(' ')
        
        result.append(c)

    return ''.join(result)


def main():
    
    with open('aaaa.txt') as fin:
        data = fin.read()

        newdata = []
        for line in data.split('\n'):
            if (m := re.match('^\s*fields\["([^"]+)"\]=&m.ops(.+)\s*$', line)) is not None:

                metric_help = get_help(m.group(2)).strip()
                metric_name = get_name(m.group(2)).strip()

                metric_help = metric_help.replace('G C ', 'GC ')
                metric_name = metric_name.replace('g_c_', 'gc_')

                metric_full_name = f'subsystem + "{metric_name}"'

                newdata.append(f"m[{metric_full_name}] = metricAttr{{\n\t\"{m.group(1)}\",\n\t\"{metric_name}\",\n\t\"{metric_help}\",\n}}")

            else:
                newdata.append(line.strip())

    with open('bbbb.txt', 'w') as fout:
        fout.write('\n'.join(newdata))


if __name__ == '__main__':
    main()
