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
//
// this package implements a bunch of convenience routines for working
// with slices.
//
// this file implements slice-comparison routines.
//
package slice_utils

//
// this function returns true if two int slices are equal i.e. for
// all i, x[i] == y[i] and 0 <= i < len(x)
//
func CmpIntSlice(x, y *[]int) bool {
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

//
// this function returns true if two int32 slices are equal i.e. for
// all i, x[i] == y[i] and 0 <= i < len(x)
//
func CmpInt32Slice(x, y *[]int32) bool {
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

//
// this function returns true if two int32 slices contain the same
// values but at different locations. slices are considered unequal if
// the number of elements are not equal i.e. len(x) != len(y)
//
func RelaxedCmpInt32Slice(x, y *[]int32) bool {
	if len(*x) != len(*y) {
		return false
	}

	for _, xv := range *x {
		found_match := false

		for _, yv := range *y {
			if xv == yv {
				found_match = true
			}
		}

		if !found_match {
			return false
		}
	}

	return true
}
