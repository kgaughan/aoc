#!/usr/bin/env python3

import collections
import re


def main():
    lhs = []
    rhs = []
    with open("input") as fh:
        for line in fh:
            match = re.match(r"^(?P<lhs>\d+)\s+(?P<rhs>\d+)", line)
            lhs.append(int(match["lhs"]))
            rhs.append(int(match["rhs"]))

    # Part 1:
    part1 = sum(abs(r - l) for l, r in zip(sorted(lhs), sorted(rhs)))
    print("Part 1:", part1)

    # Part 2:
    counts = collections.Counter(rhs)
    part2 = sum(n * counts.get(n, 0) for n in lhs)
    print("Part 2:", part2)


if __name__ == "__main__":
    main()
