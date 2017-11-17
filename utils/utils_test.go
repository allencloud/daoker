package utils

import "testing"

// TestCompareDockerVersion test
func TestCompareDockerVersion(t *testing.T) {
	type TestCase struct {
		base    string
		compare string
		result  bool
	}

	testCases := []TestCase{
		TestCase{
			base:    "1.10.0",
			compare: "1.11.0",
			result:  false,
		},
	}

	for _, testCase := range testCases {
		result, err := CompareDockerVersion(testCase.base, testCase.compare)
		if err != nil {
			t.Fatal("err")
		}
		if result != testCase.result {
			t.Fatal("error")
		}
	}
}
