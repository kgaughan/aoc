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
    return any(is_safe(readings[:i] + readings[i + 1 :]) for i in range(len(readings)))


def check_all(reports, check):
    return sum((1 if check(readings) else 0) for readings in reports)


def main():
    with open("input/day02.txt") as fh:
        reports = [list(map(int, line.rstrip().split(" "))) for line in fh]

    print("Part 1:", check_all(reports, is_safe))
    print("Part 2:", check_all(reports, is_safe_dampened))


if __name__ == "__main__":
    main()
