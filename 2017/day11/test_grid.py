#!/usr/bin/env python3

import unittest


import grid


class GridTest(unittest.TestCase):
    def test_simplify_unchanged(self):
        self.assertListEqual(grid.simplify(["ne", "ne", "ne"]), ["ne", "ne", "ne"])

    def test_simplify_home(self):
        self.assertListEqual(grid.simplify(["ne", "ne", "sw", "sw"]), [])

    def test_simplify_change_axis(self):
        self.assertListEqual(grid.simplify(["ne", "ne", "s", "s"]), ["se", "se"])

    def test_simplify_reduce(self):
        self.assertListEqual(
            grid.simplify(["se", "sw", "se", "sw", "sw"]),
            ["s", "s", "sw"],
        )
