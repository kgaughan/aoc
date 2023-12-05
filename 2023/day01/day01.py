#!/usr/bin/env python3

import re

DIGITS = {
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9",
}

DIGITS_RE = re.compile("|".join(DIGITS.keys()))


def convert_to_digits(line):
    return DIGITS_RE.sub(lambda m: DIGITS[m.group(0)], line)


def extract_digits(line):
    result = ""
    for ch in line:
        if ch.isdigit():
            result += ch
            break
    else:
        return 0
    for ch in reversed(line):
        if ch.isdigit():
            result += ch
            break
    return int(result)


def main():
    result_a = 0
    result_b = 0
    with open("input.txt", "r", encoding="utf-8") as fh:
        for line in fh:
            result_a += extract_digits(line)
            result_b += extract_digits(convert_to_digits(line))
    print("Part 1:", result_a)
    print("Part 2:", result_b)


if __name__ == "__main__":
    main()
