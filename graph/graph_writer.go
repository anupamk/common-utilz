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
// commonly used functions for writing graph structures
//
package graph

import (
	"fmt"
)

/*
 * unexported stuff
**/

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
// this function emits the graph structure in a format suitable for
// subsequent loading from LoadFromXXX(...) invokation
//
func graph_serializer(G GraphOps) string {
	str := ""

	// vertex and edge count
	str += fmt.Sprintf("%d\n", G.V())
	str += fmt.Sprintf("%d\n", G.E())

	// vertex-specific adjacency-list dump
	for v := int32(0); v < G.V(); v++ {
		for _, w := range G.Adj(v) {
			// for undirected graphs, don't dump both v-w,
			// and w-v edges
			if v > w {
				continue
			}

			str += fmt.Sprintf("%d %d\n", v, w)
		}
	}

	return str
}
