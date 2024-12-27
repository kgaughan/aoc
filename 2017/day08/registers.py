#!/usr/bin/env python3

import collections
import functools
import operator
import re


Instruction = collections.namedtuple(
    "Instruction",
    ["reg", "op", "amt", "cond_reg", "cond_op", "target"],
)


def parse_instruction(line):
    match = re.match(
        "([a-z]+) (inc|dec) (-?[0-9]+) " "if ([a-z]+) ([<>]=?|[!=]=) (-?[0-9]+)",
        line,
    )
    # Not bothering to look for parse errors.
    return Instruction(
        reg=match[1],
        op=match[2],
        amt=int(match[3]),
        cond_reg=match[4],
        cond_op=match[5],
        target=int(match[6]),
    )


def parse_file(fh):
    return [parse_instruction(line) for line in fh]


OPS = {
    "==": operator.eq,
    "!=": operator.ne,
    ">=": operator.ge,
    "<=": operator.le,
    ">": operator.gt,
    "<": operator.lt,
}


def process_instructions(instructions):
    largest = 0
    memory = collections.defaultdict(int)
    for inst in instructions:
        if OPS[inst.cond_op](memory[inst.cond_reg], inst.target):
            if inst.op == "inc":
                memory[inst.reg] += inst.amt
            elif inst.op == "dec":
                memory[inst.reg] -= inst.amt
            largest = max(largest, memory[inst.reg])
    return memory, largest


def largest_register_value(memory):
    return functools.reduce(max, memory.values())


def main():
    with open("input.txt") as fh:
        instructions = parse_file(fh)
    memory, largest_ever = process_instructions(instructions)
    print("Largest at end:", largest_register_value(memory))
    print("Largest ever:", largest_ever)


if __name__ == "__main__":
    main()
