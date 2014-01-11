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
// implement some convenience routines for parsing input string as a
// bunch of primitive types
//
package line_parser

import (
	"bufio"
	"fmt"
	"strings"
)

//
// parse an input-string into a slice of int32 values, and return
// index of the first non-int32 field
//
func ParseStringAsInt32s(line string) (retval []int32, idx int) {
	var val int32

	tokens := strings.Fields(line)
	tmp_retval := make([]int32, len(tokens))
	idx = len(tokens)

	for i, t := range tokens {
		if _, err := fmt.Sscanf(t, "%d", &val); err != nil || val < 0 {
			idx = i
			break
		}
		tmp_retval[i] = val
	}

	// compacted result-set
	retval = make([]int32, idx)
	copy(retval, tmp_retval[:idx])

	return
}

//
// return next non-empty / non-commented line from a reader
//
func GetNextLineFromReader(r *bufio.Reader, comment_char byte) (str string, err error) {
	for {
		if str, err = r.ReadString('\n'); err != nil {
			break
		}
		str = strings.TrimSpace(str)

		if len(str) != 0 && str[0] != comment_char {
			break
		}
	}

	return
}

//
// return the result of parsing next non-empty/non-commented line from
// a reader
//
func Int32sFromReader(r *bufio.Reader, comment_char byte) (retval []int32, err error) {
	var next_line string

	next_line, err = GetNextLineFromReader(r, comment_char)
	if err != nil {
		return
	}
	retval, _ = ParseStringAsInt32s(next_line)

	return
}
