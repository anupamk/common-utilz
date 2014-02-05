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
//
// commonly non-exported functions used in processing graphs.
//
package graph

import (
	"bufio"
	"fmt"
	"github.com/anupamk/common-utilz/line_parser"
	"github.com/anupamk/common-utilz/slice_utils"
	"sort"
)

func average_degree(G GraphOps) float64 {
	return float64(2 * G.V() / G.E())
}

func maximum_degree(G GraphOps) (max_degree int32) {
	max_degree = 0

	for v := int32(0); v < G.V(); v++ {
		vadj_len := int32(len(G.Adj(v)))
		if vadj_len > max_degree {
			max_degree = vadj_len
		}
	}

	return
}

//
// dump pretty-printing-string representation of the graph. output
// format is as following
//
//     <line-001> V vertices, E edges
//     <line-002> vertex-1 : adj-list-of(vertex-1)
//     <line-003> vertex-2 : adj-list-of(vertex-2)
//     <line-004> vertex-3 : adj-list-of(vertex-3)
//
func graph_stringifier(G GraphOps) string {
	str := fmt.Sprintf("%d vertices, %d edges\n", G.V(), G.E())

	for v := int32(0); v < G.V(); v++ {
		for _, w := range G.Adj(v) {
			str += fmt.Sprintf("%d ", w)
		}
		str += fmt.Sprintf("\n")
	}

	return str
}

//
// parse the input data file which is expected to be in the following
// format:
//     <line-001> number-of-vertices (V)
//     <line-002> number-of-edges (E)
//     <line-003> vertex-i vertex-j (edges[][2])
//     ....................
//     ....................
//     <line-N>   vertex-i vertex-j (edges[][2])
//
// commented-lines (starting with '#') / non-empty lines are ignored
//
func parse_graph_datafile(in *bufio.Reader) (V int32, E int32, Edges [][2]int32, err error) {
	var values []int32
	var comment_char byte = '#'

	if in == nil {
		return
	}

	// parse vertex-count
	values, err = line_parser.Int32sFromReader(in, comment_char)
	if err != nil {
		return
	}
	V = values[0]

	// parse edge-count
	values, err = line_parser.Int32sFromReader(in, comment_char)
	if err != nil {
		return
	}
	E = values[0]

	// parse edge-list
	Edges = make([][2]int32, E)
	for i := int32(0); i < E; i++ {
		values, err = line_parser.Int32sFromReader(in, comment_char)
		if err != nil {
			break
		}
		Edges[i][0], Edges[i][1] = values[0], values[1]
	}

	return
}

//
// this function returns true if the two graphs 'X' and 'Y' are
// isomorphic. the test done here is very very naive, and is more or
// less rather useless in any other setting but here.
//
func cmp_graph(X GraphOps, Y GraphOps) bool {
	// basic checks, vertex + edge count must match
	if (X.V() != Y.V()) || (X.E() != Y.E()) {
		return false
	}

	// compare adjacency list of each vertex in each graph
	for xv := int32(0); xv < X.V(); xv++ {
		adjx := X.Adj(xv)
		adjy := Y.Adj(xv)

		if cmp_adj_list(&adjx, &adjy) == false {
			return false
		}
	}

	return true
}

//
// i hope we don't run this abomination on graphs with > 2b
// vertices...
//
func cmp_adj_list(x *[]int32, y *[]int32) bool {
	if len(*x) != len(*y) {
		return false
	}

	xint := make([]int, len(*x))
	for i, xv := range *x {
		xint[i] = int(xv)
	}
	sort.Ints(xint)

	yint := make([]int, len(*y))
	for i, yv := range *y {
		yint[i] = int(yv)
	}
	sort.Ints(yint)

	return slice_utils.CmpIntSlice(&xint, &yint)
}
