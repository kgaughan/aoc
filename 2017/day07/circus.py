#!/usr/bin/env python3

import collections
import re


Program = collections.namedtuple("Program", ["name", "weight", "children"])


def parse_line(line):
    match = re.match(r"([a-z]+) \(([0-9]+)\)(?: -> (.*))?", line)
    if match is None:
        raise ValueError(line)
    if match[3] is not None:
        children = tuple(child.strip() for child in match[3].split(","))
    else:
        children = ()
    return Program(match[1], int(match[2]), children)


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


def get_weights(program, programs):
    child_weights = collections.defaultdict(list)
    total_weight = program.weight
    for child_name in program.children:
        status, name, weight = get_weights(programs[child_name], programs)
        if not status:
            return (status, name, weight)
        child_weights[weight].append(programs[child_name])
        total_weight += weight

    if len(child_weights) < 2:
        return (True, program.name, total_weight)

    # This code is written with the assumption that there is at least three
    # children and only one is unbalanced. If there's only two or all are
    # different, then finding the unbalanced one is more difficult and
    # potentially requires some backtracking.
    good_weight = None
    bad_weight = None
    bad_child = None
    for weight, children in child_weights.items():
        if len(children) > 1:
            good_weight = weight
        else:
            bad_child = children[0]
            bad_weight = weight
    return (False, bad_child.name, bad_child.weight - bad_weight + good_weight)


def find_unbalanced(root, programs):
    return get_weights(programs[root], programs)


def take_first(iterable):
    return next(iter(iterable))


def main():
    with open("input.txt") as fh:
        programs = parse_file(fh)
    parentage = build_parentage(programs)
    # Any arbitrary starting point will do.
    base = find_base(take_first(programs), parentage)
    print("Base:", base)
    _, name, weight = find_unbalanced(base, programs)
    print("Bad program is", name, "and needs to weigh", weight)


if __name__ == "__main__":
    main()
