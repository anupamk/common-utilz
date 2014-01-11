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
