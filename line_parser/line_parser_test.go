package line_parser

import (
	"bufio"
	"io"
	"strings"
	"testing"
)

//
// this function returns true if two int32 slices are equal i.e. for
// all i, x[i] == y[i] and 0 <= i < len(x)
//
func cmp_int32_slice(x, y *[]int32) bool {
	if len(*x) != len(*y) {
		return false
	}

	for i, xv := range *x {
		if xv != (*y)[i] {
			return false
		}
	}

	return true
}

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
		if idx != len(result) || cmp_int32_slice(&result, &test_val.exp_result) == false {
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

		if cmp_int32_slice(&values, &exp_value) == false {
			t.Logf("index: %d, expected: %v, got: %v\n", i, &expected_values[i], values)
			t.Fail()
		}
	}
}
