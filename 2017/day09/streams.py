#!/usr/bin/env python3

STREAM = 0
JUNK = 1
SKIP = 2


def process(stream):
    depth = 0
    score = 0
    mode = STREAM
    for ch in stream:
        if mode == SKIP:
            mode = JUNK
        elif mode == STREAM:
            if ch == "{":
                depth += 1
            elif ch == "}":
                score += depth
                depth -= 1
            elif ch == "<":
                mode = JUNK
            # Just ignoring any other characters, which only out to be ','
            # anyway.
        elif mode == JUNK:
            if ch == "!":
                mode = SKIP
            elif ch == ">":
                mode = STREAM
    return score


def count_garbage(stream):
    score = 0
    mode = STREAM
    for ch in stream:
        if mode == SKIP:
            mode = JUNK
        elif mode == STREAM:
            if ch == "<":
                mode = JUNK
        elif mode == JUNK:
            if ch == "!":
                mode = SKIP
            elif ch == ">":
                mode = STREAM
            else:
                score += 1
    return score


def main():
    with open("input.txt") as fh:
        stream = fh.read()

    print("Score:", process(stream))
    print("Garbage score:", count_garbage(stream))


if __name__ == "__main__":
    main()
