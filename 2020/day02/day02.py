#!/usr/bin/env python3

import argparse


def parse(fh):
    return []


def part_a():
    return None


def part_b():
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
