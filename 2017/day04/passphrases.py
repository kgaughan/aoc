#!/usr/bin/env python3


def slurp(path):
    result = []
    with open(path) as fh:
        for line in fh:
            result.append(line.rstrip())
    return result


def has_no_dupe(passphrase):
    words = set()
    for word in passphrase.split(' '):
        if word in words:
            return False
        words.add(word)
    return True


def has_no_anagram_dupe(passphrase):
    words = set()
    for word in passphrase.split(' '):
        word = ''.join(sorted(word))
        if word in words:
            return False
        words.add(word)
    return True


def runner(lines, fn):
    result = 0
    for line in lines:
        if fn(line):
            result += 1
    return result


def main():
    lines = slurp('input.txt')
    print(runner(lines, has_no_dupe))
    print(runner(lines, has_no_anagram_dupe))


if __name__ == '__main__':
    main()
