import unittest
import src.matrix as matrix

class TestMatrixMultiplication(unittest.TestCase):
    def setUp(self):
        self.m1 = [[1, 2, 3], [4, 5, 6]]
        self.m2 = [[7, 8], [9, 10], [11, 12]]
        self.expected_result = [[58, 64], [139, 154]]
        
        self.incompatible_m1 = [[1, 2], [3, 4]]
        self.incompatible_m2 = [[5, 6, 7]]
    
    def test_standard_mult(self):
        result = matrix.standard_mult(self.m1, self.m2)
        self.assertIsNotNone(result) 
        self.assertEqual(result, self.expected_result)

    def test_empty_rows(self):
        result = matrix.standard_mult([], [])
        self.assertIsNone(result) 

    def test_empty_cols(self):
        result = matrix.standard_mult([[]], [[]])
        self.assertIsNone(result) 
    
    def test_winograd_mult(self):
        result = matrix.winograd_mult(self.m1, self.m2)
        self.assertIsNotNone(result) 
        self.assertEqual(result, self.expected_result)
    
    def test_winograd_optimized_mult(self):
        result = matrix.winograd_optimized_mult(self.m1, self.m2)
        self.assertIsNotNone(result)
        self.assertEqual(result, self.expected_result)
    
    def test_incompatible_matrices_standard(self):
        result = matrix.standard_mult(self.incompatible_m1, self.incompatible_m2)
        self.assertIsNone(result)
    
    def test_incompatible_matrices_winograd(self):
        result = matrix.winograd_mult(self.incompatible_m1, self.incompatible_m2)
        self.assertIsNone(result)
    
    def test_incompatible_matrices_winograd_optimized(self):
        result = matrix.winograd_optimized_mult(self.incompatible_m1, self.incompatible_m2)
        self.assertIsNone(result)
    
    def test_all_methods_consistency(self):
        standard_result = matrix.standard_mult(self.m1, self.m2)
        winograd_result = matrix.winograd_mult(self.m1, self.m2)
        winograd_optimized_result = matrix.winograd_optimized_mult(self.m1, self.m2)
        
        self.assertEqual(standard_result, winograd_result)
        self.assertEqual(standard_result, winograd_optimized_result)
        self.assertEqual(winograd_result, winograd_optimized_result)
