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
// couple of commonly used algorithms for graphs are provided in this
// package.
//
// this file provides imeplentation of single-source bfs and dfs paths
// in a graph, and some routines to query/enumerate such paths
//
// this file provides the test+benchmark routines
package algorithms

import (
	"fmt"
	"github.com/anupamk/common-utilz/graph"
	"testing"
)

func ExampleDFSPath() {
	fname := "../data/graph-003.data"
	g, _ := graph.LoadGraphFromFile(fname)
	source_vertex := int32(0)

	dfs_path := DFSPath(g, source_vertex)
	fmt.Printf("%s\n", dfs_path)

	// Output:
	// source-vertex: 0
	// 0: 0
	// 1: 2
	// 2: 3
	// 3: 5
	// 4: 2
	// 5: 0
}

func ExampleBFSPath() {
	fname := "../data/graph-003.data"
	g, _ := graph.LoadGraphFromFile(fname)
	source_vertex := int32(0)

	bfs_path := BFSPath(g, source_vertex)
	fmt.Printf("%s\n", bfs_path)

	// Output:
	// source-vertex: 0
	// 0: 0
	// 1: 0
	// 2: 0
	// 3: 2
	// 4: 2
	// 5: 0
}

func ExampleBFSPathTo() {
	fname := "../data/graph-003.data"
	g, _ := graph.LoadGraphFromFile(fname)
	source_vertex := int32(0)

	bfs_path := BFSPath(g, source_vertex)

	for v := int32(0); v < g.V(); v++ {
		path, _ := bfs_path.PathTo(v)
		fmt.Printf("%d: %v\n", v, path)
	}

	// Output:
	// 0: [0]
	// 1: [0 1]
	// 2: [0 2]
	// 3: [0 2 3]
	// 4: [0 2 4]
	// 5: [0 5]
}

func ExampleDFSPathTo() {
	fname := "../data/graph-003.data"
	g, _ := graph.LoadGraphFromFile(fname)
	source_vertex := int32(0)

	bfs_path := DFSPath(g, source_vertex)

	for v := int32(0); v < g.V(); v++ {
		path, _ := bfs_path.PathTo(v)
		fmt.Printf("%d: %v\n", v, path)
	}

	// Output:
	// 0: [0]
	// 1: [0 5 3 2 1]
	// 2: [0 5 3 2]
	// 3: [0 5 3]
	// 4: [0 5 3 2 4]
	// 5: [0 5]
}

// some benchmarks

//
// all paths to all the vertices from a given source vertex in bfs
// order
//
func BenchmarkBFSPathTo(bench *testing.B) {
	fname := "../data/graph-003.data"
	g, _ := graph.LoadGraphFromFile(fname)
	source_vertex := int32(0)

	bfs_path := BFSPath(g, source_vertex)
	bench.ResetTimer()

	for i := 0; i < bench.N; i++ {
		for v := int32(0); v < g.V(); v++ {
			_, _ = bfs_path.PathTo(v)
		}
	}
}

//
// all paths to all vertices from a given source in dfs order
//
func BenchmarkDFSPathTo(bench *testing.B) {
	fname := "../data/graph-003.data"
	g, _ := graph.LoadGraphFromFile(fname)
	source_vertex := int32(0)

	bfs_path := DFSPath(g, source_vertex)
	bench.ResetTimer()

	for i := 0; i < bench.N; i++ {
		for v := int32(0); v < g.V(); v++ {
			_, _ = bfs_path.PathTo(v)
		}
	}
}
