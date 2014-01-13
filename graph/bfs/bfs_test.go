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
package bfs

import (
	"github.com/anupamk/common-utilz/graph"
	"testing"
)

var graph_bfs_data = []struct {
	fname              string
	source_vertex      int32
	connected_vertices []int32
}{
	{"../data/graph-001.data", 0, []int32{0, 1, 2, 3, 4, 5, 6}},
	{"../data/graph-001.data", 9, []int32{9, 10, 11, 12}},
	{"../data/graph-001.data", 7, []int32{7, 8}},
	{"../data/graph-003.data", 0, []int32{0, 1, 2, 3, 4, 5}},
}

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

func TestBFSSearch(t *testing.T) {
	for _, graph_data := range graph_bfs_data {
		g, _ := graph.LoadGraphFromFile(graph_data.fname)
		bfs_g := New(g, graph_data.source_vertex)
		cv := bfs_g.ConnectedVertices()

		if cmp_int32_slice(&graph_data.connected_vertices, &cv) != true {
			t.Logf("expected: %v, got: %v\n", graph_data.connected_vertices, cv)
			t.Fail()
		}
	}
}

// benchmark

//
// pretty rudimentary actually, for all the vertices in a graph,
// create bfs with the given vertex as source.
//
func BenchmarkBreadthFirstSearch(bench *testing.B) {
	graph_fname := "../data/graph-003.data"
	g, _ := graph.LoadGraphFromFile(graph_fname)

	for i := 0; i < bench.N; i++ {
		for v := int32(0); v < g.V(); v++ {
			tmp := New(g, v)
			if tmp.Count() == 0 {
				panic("Ooops")
			}
		}
	}
}
