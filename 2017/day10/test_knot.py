#!/usr/bin/env python3

import unittest

import knot


class KnotTest(unittest.TestCase):

    def test_tie(self):
        self.assertListEqual(knot.part1([3, 4, 1, 5], size=5),
                             [3, 4, 2, 1, 0])

    def test_hash(self):
        fixtures = {
            '': 'a2582a3a0e66e6e86e3812dcb672a272',
            'AoC 2017': '33efeb34ea91902bb2f59c9920caa6cd',
            '1,2,3': '3efbe78a8d82f29979031a4aa0b16a9d',
            '1,2,4': '63960835bcdc130f0b66d7ff4f6a5a8e',
        }
        for src, expected in fixtures.items():
            self.assertEqual(knot.part2([ord(ch) for ch in src]), expected)
