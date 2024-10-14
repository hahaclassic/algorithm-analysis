import unittest

from src import levenshtein as lvnsht

class TestLevenshtein(unittest.TestCase):
    def testRecursiveLevenshtein(self):
        self.assertEqual(lvnsht.RecursiveLevenshtein("Hello world", "ello kordH"), 4)

    def testRecursiveCacheLevenshtein(self):
        self.assertEqual(lvnsht.RecursiveCacheLevenshtein("Hello world", "ello kordH"), 4)

    def testDynamicLevenshtein(self):
        self.assertEqual(lvnsht.DynamicLevenshtein("Hello world", "ello kordH"), 4)

    def testDynamicDamerauLevenshtein(self):
        self.assertEqual(lvnsht.DynamicDamerauLevenshtein("Hello world", "ello kordH"), 3)
