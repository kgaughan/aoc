#!/usr/bin/env python3

import collections
import re


Program = collections.namedtuple('Program', ['name', 'weight', 'children'])


def parse_line(line):
    match = re.match(r'([a-z]+) \(([0-9]+)\)(?: -> (.*))?', line)
    if match is None:
        raise ValueError(line)
    if match.group(3) is not None:
        children = tuple(child.strip()
                         for child in match.group(3).split(','))
    else:
        children = ()
    return Program(match.group(1), int(match.group(2)), children)
