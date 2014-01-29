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
// this file implements testing stuff for slice-comparison routines
package slice_utils

import (
	"testing"
)

func TestCmpIntSlice(t *testing.T) {
	var int_slice_test = []struct {
		x    []int
		y    []int
		want bool
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{1, 3, 3}, []int{1, 2, 3}, false},
		{[]int{}, []int{}, true},
		{[]int{}, []int{123, 12, 3443}, false},
	}

	for _, test_val := range int_slice_test {
		got := CmpIntSlice(&test_val.x, &test_val.y)
		if got != test_val.want {
			t.Logf("Failed: wanted: %v, got: %v, slice-x: %v, slice-y: %v\n",
				test_val.want, got, test_val.x, test_val.y)
			t.Fail()
		}
	}
}

func TestCmpInt32Slice(t *testing.T) {
	var int32_slice_test = []struct {
		x    []int32
		y    []int32
		want bool
	}{
		{[]int32{1, 2, 3}, []int32{1, 2, 3}, true},
		{[]int32{1, 3, 3}, []int32{1, 2, 3}, false},
		{[]int32{}, []int32{}, true},
	}

	for _, test_val := range int32_slice_test {
		got := CmpInt32Slice(&test_val.x, &test_val.y)
		if got != test_val.want {
			t.Logf("Failed: wanted: %v, got: %v, slice-x: %v, slice-y: %v\n",
				test_val.want, got, test_val.x, test_val.y)
			t.Fail()
		}
	}

}

func TestRelaxedCmpInt32Slice(t *testing.T) {
	var int32_slice_test = []struct {
		x    []int32
		y    []int32
		want bool
	}{
		{[]int32{1, 2, 3}, []int32{3, 1, 2}, true},
		{[]int32{1, 3, 2, 4}, []int32{1, 2, 3}, false},
		{[]int32{1, 3, 3}, []int32{1, 4, 3}, false},
		{[]int32{}, []int32{}, true},
	}

	for _, test_val := range int32_slice_test {
		got := RelaxedCmpInt32Slice(&test_val.x, &test_val.y)
		if got != test_val.want {
			t.Logf("Failed: wanted: %v, got: %v, slice-x: %v, slice-y: %v\n",
				test_val.want, got, test_val.x, test_val.y)
			t.Fail()
		}
	}

}

func TestCmpStringSlice(t *testing.T) {
	var string_slice_test = []struct {
		x    []string
		y    []string
		want bool
	}{
		{[]string{"abcd", "efg", "hijkl"}, []string{"abcd", "efg", "hijkl"}, true},
		{[]string{"abcd", "efg", "hijkl", "alpha"}, []string{"abcd", "efg", "hijkl"}, false},
		{[]string{"abcd", "efg", "hijkl"}, []string{"efg", "hijkl", "abcd"}, false},
		{[]string{}, []string{}, true},
	}

	for _, test_val := range int32_slice_test {
		got := CmpStringSlice(&test_val.x, &test_val.y)
		if got != test_val.want {
			t.Logf("Failed: wanted: %v, got: %v, slice-x: %v, slice-y: %v\n",
				test_val.want, got, test_val.x, test_val.y)
			t.Fail()
		}
	}

}

func TestRelaxedCmpStringSlice(t *testing.T) {
	var string_slice_test = []struct {
		x    []string
		y    []string
		want bool
	}{
		{[]string{"abcd", "efg", "hijkl"}, []string{"efg", "hijkl", "abcd"}, true},
		{[]string{"abcd", "efg", "hijkl", "alpha"}, []string{"abcd", "efg", "hijkl"}, false},
		{[]string{"abcd", "efg", "efg"}, []string{"efg", "hijkl", "abcd"}, false},
		{[]string{}, []string{}, true},
	}

	for i, test_val := range int32_slice_test {
		got := RelaxedCmpStringSlice(&test_val.x, &test_val.y)
		if got != test_val.want {
			t.Logf("Failed-%d: wanted: %v, got: %v, slice-x: %v, slice-y: %v\n",
				i, test_val.want, got, test_val.x, test_val.y)
			t.Fail()
		}
	}

}
