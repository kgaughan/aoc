#!/usr/bin/env python3

import argparse
import itertools


def parse(fh):
    return [int(arg.strip()) for arg in fh]


def part_a(numbers):
    for fst, snd in itertools.combinations(numbers, 2):
        if fst + snd == 2020:
            return fst * snd
    return None


def part_b(numbers):
    for fst, snd, trd in itertools.combinations(numbers, 3):
        if fst + snd + trd == 2020:
            return fst * snd * trd
    return None


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("infile", type=argparse.FileType())
    args = parser.parse_args()
    records = parse(args.infile)
    print("a:", part_a(records))
    print("b:", part_b(records))


if __name__ == "__main__":
    main()
