#!/usr/bin/env python3

import argparse
import itertools
import sys

def day01a(numbers):
    for fst, snd in itertools.combinations(numbers, 2):
        if fst + snd == 2020:
            return fst * snd
    return None


def day01b(numbers):
    for fst, snd, trd in itertools.combinations(numbers, 3):
        if fst + snd + trd == 2020:
            return fst * snd * trd
    return None


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("infile", type=argparse.FileType())
    args = parser.parse_args()
    numbers = [int(arg.strip()) for arg in args.infile]
    print("day01a:", day01a(numbers))
    print("day01b:", day01b(numbers))


if __name__ == "__main__":
    main()
