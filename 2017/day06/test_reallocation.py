#!/usr/bin/env python3

import unittest

import reallocation


class ReallocationTest(unittest.TestCase):
    def test(self):
        self.assertEqual(reallocation.find_cycle([0, 2, 7, 0]), (5, 4))
