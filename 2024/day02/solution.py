#!/usr/bin/env python3

import itertools


def diffs(readings):
    return [b - a for a, b in itertools.pairwise(readings)]


def is_safe(readings):
    sign_sum = 0
    for diff in diffs(readings):
        if diff == 0 or abs(diff) > 3:
            return False
        sign_sum += -1 if diff < 0 else 1
    return abs(sign_sum) == len(readings) - 1


def is_safe_dampened(readings):
    if is_safe(readings):
        return True
    # Ugly bruteforcing...
    for i in range(len(readings)):
        if is_safe([n for j, n in enumerate(readings) if i != j]):
            return True
    return False



def main():
    with open("input") as fh:
        reports = [list(map(int, line.rstrip().split(" "))) for line in fh]

    part1 = sum((1 if is_safe(readings) else 0) for readings in reports)
    print("Part 1:", part1)

    part2 = sum((1 if is_safe_dampened(readings) else 0) for readings in reports)
    print("Part 2:", part2)


if __name__ == "__main__":
    main()
