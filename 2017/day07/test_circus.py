#!/usr/bin/env python3

import io
import unittest

import circus


FIXTURE = """\
pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)
"""


class CircusTest(unittest.TestCase):
    def test_parse_fail(self):
        with self.assertRaises(ValueError, msg="ktlj"):
            circus.parse_line("ktlj")

    def test_parse_no_children(self):
        result = circus.parse_line("ktlj (57)")
        self.assertEqual(result.name, "ktlj")
        self.assertEqual(result.weight, 57)
        self.assertTupleEqual(result.children, ())

    def test_parse_with_children(self):
        result = circus.parse_line("fwft (72) -> ktlj, cntj, xhth")
        self.assertEqual(result.name, "fwft")
        self.assertEqual(result.weight, 72)
        self.assertTupleEqual(result.children, ("ktlj", "cntj", "xhth"))

    def test_parse_file(self):
        with io.StringIO(FIXTURE) as fh:
            programs = circus.parse_file(fh)
        self.assertEqual(len(programs), 13)

    def test_find_base(self):
        with io.StringIO(FIXTURE) as fh:
            programs = circus.parse_file(fh)
        parentage = circus.build_parentage(programs)
        base = circus.find_base(circus.take_first(programs), parentage)
        self.assertEqual(base, "tknk")

    def test_find_unbalanced(self):
        with io.StringIO(FIXTURE) as fh:
            programs = circus.parse_file(fh)
        parentage = circus.build_parentage(programs)
        base = circus.find_base(circus.take_first(programs), parentage)
        status, name, good_weight = circus.find_unbalanced(base, programs)
        self.assertFalse(status)
        self.assertEqual(name, "ugml")
        self.assertEqual(good_weight, 60)
