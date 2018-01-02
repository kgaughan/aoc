#!/usr/bin/env python3

import io
import unittest

import registers


FIXTURE = """\
b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10
"""


class RegistersTest(unittest.TestCase):

    def test_parse_instruction(self):
        inst = registers.parse_instruction('b inc 5 if a > 1')
        self.assertEqual(inst.reg, 'b')
        self.assertEqual(inst.op, 'inc')
        self.assertEqual(inst.amt, 5)
        self.assertEqual(inst.cond_reg, 'a')
        self.assertEqual(inst.cond_op, '>')
        self.assertEqual(inst.target, 1)

    def test_part1(self):
        with io.StringIO(FIXTURE) as fh:
            instructions = registers.parse_file(fh)
        memory = registers.process_instructions(instructions)
        self.assertEqual(registers.largest_register_value(memory), 1)
        self.assertListEqual(sorted(memory.keys()), ['a', 'b', 'c'])
        self.assertEqual(memory['a'], 1)
        self.assertEqual(memory['b'], 0)
        self.assertEqual(memory['c'], -10)
