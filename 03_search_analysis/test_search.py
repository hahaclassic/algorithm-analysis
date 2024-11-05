import unittest

from src import list_funcs as lf

class TestSearch(unittest.TestCase):
    def setUp(self) -> None:
        self.sorted_array = [-1, -2, 3, 4, 5, 6, 7, 8, 9, 10]
        self.array = [1, -10, 3, 4, 8, -6, 2, 7, 9, 8, -5]

    def testLinearSearchFoundSorted(self) -> None:
        self.assertEqual(lf.linearSearch(self.sorted_array, 7), (6, 7))

    def testLinearSearchFound(self) -> None:
        self.assertEqual(lf.linearSearch(self.array, 7), (7, 8))

    def testLinearSearchNotFound(self) -> None:
        self.assertEqual(lf.linearSearch(self.array, 11), (-1, 11))

    def testLinearSearchEmptyList(self) -> None:
        self.assertEqual(lf.linearSearch([], 11), (-1, 0))

    def testBinarySearchFound(self) -> None:
        self.assertEqual(lf.binarySearch(self.sorted_array, 7), (6, 4))

    def testBinarySearchNotFound(self) -> None:
        self.assertEqual(lf.binarySearch(self.sorted_array, 11), (-1, 4))

    def testBinarySearchEmptyList(self) -> None:
        self.assertEqual(lf.binarySearch([], 11), (-1, 0))
