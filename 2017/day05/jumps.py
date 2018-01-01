#!/usr/bin/env python3


def eval_jumps1(jumps):
    cloned = list(jumps)
    pc = 0
    i = 0
    while 0 <= pc < len(cloned):
        old_pc = cloned[pc]
        cloned[pc] += 1
        pc += old_pc
        i += 1
    return i


def eval_jumps2(jumps):
    cloned = list(jumps)
    pc = 0
    i = 0
    while 0 <= pc < len(cloned):
        old_pc = cloned[pc]
        if cloned[pc] >= 3:
            cloned[pc] -= 1
        else:
            cloned[pc] += 1
        pc += old_pc
        i += 1
    return i


def main():
    jumps = []
    with open('input.txt') as fh:
        for line in fh:
            jumps.append(int(line))

    print(eval_jumps1(jumps))
    print(eval_jumps2(jumps))


if __name__ == '__main__':
    main()
