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


def find_base(programs):
    base = None
    parents = {}
    for program in programs.values():
        for child in program.children:
            parents[child] = program.name
        # While we're building up the mapping, get as far down the tree as we
        # can.
        if base is None or base in program.children:
            base = program.name

    # Travel down the remainder of the tree.
    while base in parents:
        base = parents[base]

    return base


def main():
    with open('input.txt') as fh:
        programs = parse_file(fh)
    print(find_base(programs))


if __name__ == '__main__':
    main()
