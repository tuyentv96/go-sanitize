package sanitize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrimStruct(t *testing.T) {
	type testStruct struct {
		A string
		B string
		C string
	}

	testCases := []struct {
		input  testStruct
		expect testStruct
	}{
		{
			input: testStruct{
				A: " A ",
				B: " B ",
				C: " C ",
			},
			expect: testStruct{
				A: "A",
				B: "B",
				C: "C",
			},
		},
	}

	for _, testCase := range testCases {
		err := TrimSpace(&testCase.input)
		if err != nil {
			t.Fatalf("got an err: %s", err)
		}

		assert.Equal(t, testCase.expect, testCase.input)
	}

}
