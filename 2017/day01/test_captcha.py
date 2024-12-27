#!/usr/bin/env python3
import unittest

import captcha


class CaptchaTests(unittest.TestCase):
    def test_part1(self):
        self.assertEqual(captcha.captcha1("1122"), 3)
        self.assertEqual(captcha.captcha1("1111"), 4)
        self.assertEqual(captcha.captcha1("1234"), 0)
        self.assertEqual(captcha.captcha1("91212129"), 9)

    def test_part2(self):
        self.assertEqual(captcha.captcha2("1212"), 6)
        self.assertEqual(captcha.captcha2("1221"), 0)
        self.assertEqual(captcha.captcha2("123425"), 4)
        self.assertEqual(captcha.captcha2("123123"), 12)
        self.assertEqual(captcha.captcha2("12131415"), 4)
