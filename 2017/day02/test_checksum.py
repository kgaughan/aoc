#!/usr/bin/env python3
import unittest

import checksum


class ChecksumTests(unittest.TestCase):

    def test_part1(self):
        fixture = [
            [5, 1, 9, 5],
            [7, 5, 3],
            [2, 4, 6, 8],
        ]
        self.assertEqual(checksum.checksum_line(fixture[0]), 8)
        self.assertEqual(checksum.checksum_line(fixture[1]), 4)
        self.assertEqual(checksum.checksum_line(fixture[2]), 6)
        self.assertEqual(checksum.checksum1(fixture), 18)

    def test_part2(self):
        fixture = [
            [5, 9, 2, 8],
            [9, 4, 7, 3],
            [3, 8, 6, 5],
        ]
        self.assertEqual(checksum.checksum_line_evens(fixture[0]), 4)
        self.assertEqual(checksum.checksum_line_evens(fixture[1]), 3)
        self.assertEqual(checksum.checksum_line_evens(fixture[2]), 2)
        self.assertEqual(checksum.checksum2(fixture), 9)
