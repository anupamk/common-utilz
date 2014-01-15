//
// Copyright (c) 2014, Anupam Kapoor. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:

// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// Author: anupam.kapoor@gmail.com (Anupam Kapoor)
//
package line_parser

import (
	"bufio"
	"github.com/anupamk/common-utilz/slice_utils"
	"io"
	"strings"
	"testing"
)

func TestInt32Parse(t *testing.T) {
	var int32_tests = []struct {
		str        string
		exp_result []int32
	}{
		{"123", []int32{123}},
		{"1234 123", []int32{1234, 123}},
		{"#1234	907		145		111", []int32{}},
		{"456 1234 9021", []int32{456, 1234, 9021}},
		{"12345 hello 456", []int32{12345}},
	}

	for _, test_val := range int32_tests {
		result, idx := ParseStringAsInt32s(test_val.str)
		if idx != len(result) || slice_utils.CmpInt32Slice(&result, &test_val.exp_result) == false {
			t.Logf("expected: %v, got: %v\n", test_val.exp_result, result)
			t.Fail()
		}
	}
}

func TestInt32sFromReader(t *testing.T) {
	var test_string = `
1234
3456
#567 890 987
		



11 12 123 13456
		123		9087		9123

hello

#

`
	var expected_values = [...][]int32{
		[]int32{1234},
		[]int32{3456},
		[]int32{11, 12, 123, 13456},
		[]int32{123, 9087, 9123},
	}

	str_reader := strings.NewReader(test_string)
	str_bufio_reader := bufio.NewReader(str_reader)

	for i, exp_value := range expected_values {
		values, err := Int32sFromReader(str_bufio_reader, '#')
		if err != nil && err == io.EOF {
			if i != len(exp_value) {
				t.Logf("Error: missed parsing some values")
				t.Fail()
			}

			return
		}

		if slice_utils.CmpInt32Slice(&values, &exp_value) == false {
			t.Logf("index: %d, expected: %v, got: %v\n", i, &expected_values[i], values)
			t.Fail()
		}
	}
}
