#!/usr/bin/env python3

import itertools


def parse(line):
    return [int(entry) for entry in line.rstrip().split("\t")]


def parse_file(path):
    lines = []
    with open(path) as fh:
        for line in fh:
            lines.append(parse(line))
    return lines


def checksum_line(line):
    minimum = line[0]
    maximum = line[0]
    for value in line:
        if value < minimum:
            minimum = value
        if value > maximum:
            maximum = value
    return maximum - minimum


def checksum1(lines):
    return sum(checksum_line(line) for line in lines)


def sorted_combinations(line):
    return itertools.combinations(sorted(line), 2)


def checksum_line_evens(line):
    result = 0
    for denom, numer in sorted_combinations(line):
        quotient, mod = divmod(numer, denom)
        if mod == 0:
            result += quotient
    return result


def checksum2(lines):
    return sum(checksum_line_evens(line) for line in lines)


def main():
    lines = parse_file("input.txt")
    print(checksum1(lines))
    print(checksum2(lines))


if __name__ == "__main__":
    main()
