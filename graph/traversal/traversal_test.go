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
package traversal

import (
	"fmt"
	"github.com/anupamk/common-utilz/graph"
	"github.com/anupamk/common-utilz/slice_utils"
	"testing"
)

// ensure vertices are traversed bfs order
func ExampleCheckBFSTraversalOrder() {
	fname := "../data/graph-003.data"
	g, _ := graph.LoadGraphFromFile(fname)
	source_vertex := int32(0)

	bfs_walker := BFSGraphSubsetWalker(g, source_vertex)

	for v, err := bfs_walker(); err != EOGS; v, err = bfs_walker() {
		fmt.Println(v)
	}

	// Output:
	// 0
	// 2
	// 1
	// 5
	// 3
	// 4
}

// ensure vertices are traversed in dfs order
func ExampleCheckDFSSubsetTraversalOrder() {
	fname := "../data/graph-003.data"
	g, _ := graph.LoadGraphFromFile(fname)
	source_vertex := int32(0)

	dfs_walker := DFSGraphSubsetWalker(g, source_vertex)
	for v, err := dfs_walker(); err != EOGS; v, err = dfs_walker() {
		fmt.Println(v)
	}

	// Output:
	// 1
	// 4
	// 2
	// 3
	// 5
	// 0
}

// test to see if graph subset traversal visits expected set of
// vertices for a given source.
func TestCheckGraphSubsetTraversal(t *testing.T) {
	fname := "../data/graph-001.data"
	g, _ := graph.LoadGraphFromFile(fname)
	source_vertex := int32(0)
	actual_vv := &[]int32{0, 1, 2, 6, 5, 3, 4}

	visited_nodes := func(walker GraphSubsetWalker) *[]int32 {
		i, vv := 0, make([]int32, g.V())

		// visit vertices
		for v, err := walker(); err != EOGS; v, err = walker() {
			vv[i] = v
			i += 1
		}

		// compact generated list
		x := make([]int32, i)
		copy(x, vv)

		return &x
	}

	// test bfs traversal
	bfs_vv := visited_nodes(BFSGraphSubsetWalker(g, source_vertex))
	if !slice_utils.RelaxedCmpInt32Slice(bfs_vv, actual_vv) {
		t.Logf("failed\nactual-visited-vertices: %v\nexpected-visited-vertices: %v\n", bfs_vv, actual_vv)
		t.Fail()
	}

	// test dfs traversal
	dfs_vv := visited_nodes(DFSGraphSubsetWalker(g, source_vertex))
	if !slice_utils.RelaxedCmpInt32Slice(dfs_vv, actual_vv) {
		t.Logf("failed\nactual-visited-vertices: %v\nexpected-visited-vertices: %v\n", bfs_vv, actual_vv)
		t.Fail()
	}
}

// test to see if graph traversal does cover all the nodes...
func TestCheckGraphTraversal(t *testing.T) {
	fname := "../data/graph-001.data"
	g, _ := graph.LoadGraphFromFile(fname)
	actual_vv := &[]int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	visited_nodes := func(walker GraphWalker) *[]int32 {
		i, vv := 0, make([]int32, g.V())

		// visit vertices
		for v, err := walker(); err != EOG; v, err = walker() {
			vv[i] = v
			i += 1
		}

		return &vv
	}

	// test bfs
	bfs_vv := visited_nodes(BFSGraphWalker(g))
	if !slice_utils.RelaxedCmpInt32Slice(bfs_vv, actual_vv) {
		t.Logf("failed\nactual-visited-vertices: %v\nexpected-visited-vertices: %v\n", bfs_vv, actual_vv)
		t.Fail()
	}

	// test dfs
	dfs_vv := visited_nodes(BFSGraphWalker(g))
	if !slice_utils.RelaxedCmpInt32Slice(dfs_vv, actual_vv) {
		t.Logf("failed\nactual-visited-vertices: %v\nexpected-visited-vertices: %v\n", dfs_vv, actual_vv)
		t.Fail()
	}

}

// benchmark various traversals

// bfs
func BenchmarkBFSGraphTraversal(bench *testing.B) {
	fname := "../data/graph-004.data"
	g, _ := graph.LoadGraphFromFile(fname)
	bench.ResetTimer()

	for i, gw := 0, BFSGraphWalker(g); i < bench.N; i++ {
		for _, err := gw(); err != EOG; _, err = gw() {
		}
	}
}

// dfs
func BenchmarkDFSGraphTraversal(bench *testing.B) {
	fname := "../data/graph-004.data"
	g, _ := graph.LoadGraphFromFile(fname)
	bench.ResetTimer()

	for i, gw := 0, DFSGraphWalker(g); i < bench.N; i++ {
		for _, err := gw(); err != EOG; _, err = gw() {
		}
	}
}
