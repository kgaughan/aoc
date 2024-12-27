#!/usr/bin/env python3

import unittest

import passphrases


class PassphrasesTests(unittest.TestCase):
    def assert_dupe(self, phrase):
        self.assertFalse(passphrases.has_no_dupe(phrase))

    def assert_no_dupe(self, phrase):
        self.assertTrue(passphrases.has_no_dupe(phrase))

    def assert_anagram_dupe(self, phrase):
        self.assertFalse(passphrases.has_no_anagram_dupe(phrase))

    def assert_no_anagram_dupe(self, phrase):
        self.assertTrue(passphrases.has_no_anagram_dupe(phrase))

    def test_no_dupes(self):
        self.assert_no_dupe("aa bb cc dd ee")
        self.assert_dupe("aa bb cc dd aa")
        self.assert_no_dupe("aa bb cc dd aaa")

    def test_no_anagram_dupes(self):
        self.assert_no_anagram_dupe("abcde fghij")
        self.assert_anagram_dupe("abcde xyz ecdab")
        self.assert_no_anagram_dupe("a ab abc abd abf abj")
        self.assert_no_anagram_dupe("iii oiii ooii oooi oooo")
        self.assert_anagram_dupe("oiii ioii iioi iiio")
