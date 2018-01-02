#!/usr/bin/env python3

import unittest

import circus


class CircusTest(unittest.TestCase):

    def test_parse_fail(self):
        with self.assertRaises(ValueError, msg='ktlj'):
            circus.parse_line('ktlj')

    def test_parse_no_children(self):
        result = circus.parse_line('ktlj (57)')
        self.assertEqual(result.name, 'ktlj')
        self.assertEqual(result.weight, 57)
        self.assertTupleEqual(result.children, ())

    def test_parse_with_children(self):
        result = circus.parse_line('fwft (72) -> ktlj, cntj, xhth')
        self.assertEqual(result.name, 'fwft')
        self.assertEqual(result.weight, 72)
        self.assertTupleEqual(result.children, ('ktlj', 'cntj', 'xhth'))
