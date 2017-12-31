#!/usr/bin/env python3


def runner(path, fn):
    with open(path) as fh:
        for i, line in enumerate(fh):
            print('{}: {}'.format(i + 1, fn(line.rstrip())))


def captcha1(data):
    result = 0
    last = None
    for ch in data + data[0]:
        if last == ch:
            result += int(ch)
        last = ch
    return result


def captcha2(data):
    half = len(data) // 2
    return sum(2 * int(ch1)
               for ch1, ch2 in zip(data[:half], data[half:])
               if ch1 == ch2)


def part1():
    runner('input.txt', captcha1)


def part2():
    runner('input.txt', captcha2)
