#!/usr/bin/env python3

import functools
import operator


def tie(blk, lengths, pos=0, skip=0):
    size = len(blk)
    for length in lengths:
        i_end = pos + length
        for n in range(length // 2):
            i = (pos + n) % size
            j = (i_end - n - 1) % size
            blk[i], blk[j] = blk[j], blk[i]
        pos = (i_end + skip) % size
        skip += 1
    return blk, pos, skip


def part1(lengths, size=256):
    blk, _, _ = tie(list(range(size)), lengths)
    return blk


def part2(lengths):
    pos = 0
    skip = 0
    blk = list(range(256))
    # Note: This does a join, not an append.
    lengths = lengths + [17, 31, 73, 47, 23]
    # Apply rounds.
    for _ in range(64):
        blk, pos, skip = tie(blk, lengths, pos, skip)
    # Reduce the sparse hash down to the dense hash.
    dense = [functools.reduce(operator.xor, blk[i:i+16], 0)
             for i in range(0, 256, 16)]
    # Convert to hex.
    return ''.join('{:02x}'.format(n) for n in dense)


def main():
    with open('input.txt') as fh:
        lengths = [int(length) for length in fh.readline().split(',')]
    blk = part1(lengths)
    print('Product:', blk[0] * blk[1])

    with open('input.txt') as fh:
        lengths = [ord(ch) for ch in fh.readline().rstrip()]
    print(part2(lengths))


if __name__ == '__main__':
    main()
