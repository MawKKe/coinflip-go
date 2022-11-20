package main

import (
	"testing"
)

func TestBitToString(t *testing.T) {
	type testData struct {
		bit    int
		yesno  bool
		expect string
	}
	data := []testData{
		testData{
			bit:    0,
			yesno:  false,
			expect: "tails",
		},
		testData{
			bit:    1,
			yesno:  false,
			expect: "heads",
		},
		testData{
			bit:    2,
			yesno:  false,
			expect: "tails",
		},
		testData{
			bit:    3,
			yesno:  false,
			expect: "heads",
		},
		testData{
			bit:    0,
			yesno:  true,
			expect: "no",
		},
		testData{
			bit:    1,
			yesno:  true,
			expect: "yes",
		},
		testData{
			bit:    2,
			yesno:  true,
			expect: "no",
		},
		testData{
			bit:    3,
			yesno:  true,
			expect: "yes",
		},
	}
	for _, elem := range data {
		res := bitToString(elem.bit, elem.yesno)
		if res != elem.expect {
			t.Fatalf("bitToString(%d, %t) = %q, expected %q",
				elem.bit, elem.yesno, res, elem.expect)
		}
	}
}
