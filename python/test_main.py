import unittest

class TestMain(unittest.TestCase):
    def setUp(self):
        pass
    
    def test_passing(self):
        dummy_result = 8
        self.assertEqual(dummy_result, 8)
    
    def tearDown(self):
        pass