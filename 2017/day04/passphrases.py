#!/usr/bin/env python3


def slurp(path):
    result = []
    with open(path) as fh:
        result.extend(line.rstrip() for line in fh)
    return result


def has_no_dupe(passphrase):
    words = set()
    for word in passphrase.split(" "):
        if word in words:
            return False
        words.add(word)
    return True


def has_no_anagram_dupe(passphrase):
    words = set()
    for word in passphrase.split(" "):
        word = "".join(sorted(word))
        if word in words:
            return False
        words.add(word)
    return True


def runner(lines, fn):
    return sum(bool(fn(line)) for line in lines)


def main():
    lines = slurp("input.txt")
    print(runner(lines, has_no_dupe))
    print(runner(lines, has_no_anagram_dupe))


if __name__ == "__main__":
    main()
