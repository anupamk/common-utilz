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
package dfs

import (
	"fmt"
	"github.com/anupamk/common-utilz/graph"
	"github.com/anupamk/common-utilz/slice_utils"
	"testing"
)

var graph_dfs_data = []struct {
	fname              string
	source_vertex      int32
	connected_vertices []int32
}{
	{"../data/graph-001.data", 0, []int32{0, 1, 2, 3, 4, 5, 6}},
	{"../data/graph-001.data", 9, []int32{9, 10, 11, 12}},
	{"../data/graph-001.data", 7, []int32{7, 8}},
	{"../data/graph-003.data", 0, []int32{0, 1, 2, 3, 4, 5}},
}

func TestDFSSearch(t *testing.T) {
	for _, graph_data := range graph_dfs_data {
		g, _ := graph.LoadFromFile(graph_data.fname)
		dfs_g := New(g, graph_data.source_vertex)
		cv := dfs_g.ConnectedVertices()

		if slice_utils.CmpInt32Slice(&graph_data.connected_vertices, &cv) != true {
			t.Logf("expected: %v, got: %v\n", graph_data.connected_vertices, cv)
			t.Fail()
		}
	}
}

func ExampleDepthFirstPath() {
	graph_fname := "../data/graph-003.data"
	source_vertex := int32(0)
	g, _ := graph.LoadFromFile(graph_fname)

	dfs_g := New(g, source_vertex)
	for dst_vertex := source_vertex; dst_vertex < g.V(); dst_vertex++ {
		dst_path := dfs_g.Path(dst_vertex)
		fmt.Printf("from %d to %d: %v\n", dfs_g.source, dst_vertex, dst_path)
	}

	// Output:
	// from 0 to 0: [0]
	// from 0 to 1: [0 2 1]
	// from 0 to 2: [0 2]
	// from 0 to 3: [0 2 3]
	// from 0 to 4: [0 2 3 4]
	// from 0 to 5: [0 2 3 5]

	return
}

// benchmark

//
// pretty rudimentary actually, for all the vertices in a graph,
// create dfs with the given vertex as source.
//
func BenchmarkDepthFirstSearch(bench *testing.B) {
	graph_fname := "../data/graph-003.data"
	g, _ := graph.LoadFromFile(graph_fname)

	for i := 0; i < bench.N; i++ {
		for v := int32(0); v < g.V(); v++ {
			tmp := New(g, v)
			if tmp.Count() == 0 {
				panic("Ooops")
			}
		}
	}
}
