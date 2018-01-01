#!/usr/bin/env python3

import unittest

import jumps


class TestJumps(unittest.TestCase):

    def test_part1(self):
        self.assertEqual(jumps.eval_jumps1([0, 3, 0, 1, -3]), 5)

    def test_part2(self):
        self.assertEqual(jumps.eval_jumps2([0, 3, 0, 1, -3]), 10)
