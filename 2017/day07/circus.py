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


def parse_file(fh):
    programs = {}
    for line in fh:
        program = parse_line(line)
        programs[program.name] = program
    return programs


def build_parentage(programs):
    parentage = {}
    for program in programs.values():
        for child in program.children:
            parentage[child] = program
    return parentage


def find_base(name, parentage):
    base = name
    while base in parentage:
        base = parentage[base].name
    return base


def take_first(iterable):
    return next(iter(iterable))


def main():
    with open('input.txt') as fh:
        programs = parse_file(fh)
    parentage = build_parentage(programs)
    # Any arbitrary starting point will do.
    base = find_base(take_first(programs), parentage)
    print('Base:', base)


if __name__ == '__main__':
    main()
