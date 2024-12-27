#!/usr/bin/env python3

import unittest

import streams


class StreamsTest(unittest.TestCase):
    def test_stream(self):
        self.assertEqual(streams.process("{}"), 1)
        self.assertEqual(streams.process("{{{}}}"), 6)
        self.assertEqual(streams.process("{{},{}}"), 5)
        self.assertEqual(streams.process("{{{},{},{{}}}}"), 16)
        self.assertEqual(streams.process("{<{},{},{{}}>}"), 1)
        self.assertEqual(streams.process("{<a>,<a>,<a>,<a>}"), 1)
        self.assertEqual(streams.process("{{<ab>},{<ab>},{<ab>},{<ab>}}"), 9)
        self.assertEqual(streams.process("{{<!!>},{<!!>},{<!!>},{<!1>}}"), 9)
        self.assertEqual(streams.process("{{<a!>},{<a!>},{<a!>},{<ab>}}"), 3)

    def test_garbage(self):
        self.assertEqual(streams.count_garbage("<>"), 0)
        self.assertEqual(streams.count_garbage("<random characters>"), 17)
        self.assertEqual(streams.count_garbage("<<<<>"), 3)
        self.assertEqual(streams.count_garbage("<{!>}>"), 2)
        self.assertEqual(streams.count_garbage("<!!>"), 0)
        self.assertEqual(streams.count_garbage("<!!!>>"), 0)
        self.assertEqual(streams.count_garbage('<{o"i!a,<{i<a>'), 10)
