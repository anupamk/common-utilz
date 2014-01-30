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
// this package implements the symbol graph where vertex names are
// strings and number of edges/vertices are implicitly defined. this
// is more typical of real-world (tm) graph applications
//
// this file implements the testing routine for symbol-graphs
//
package symbol_graph

import (
	"github.com/anupamk/common-utilz/slice_utils"
	"testing"
)

var symbol_graph_data_files = [...]string{
	"../data/symgraph-001.data",
}

//
// load a symbol graph from a file, and compare the adjacency list of
// the graph with the expected adjacency
//
func TestNumericAdj(t *testing.T) {
	fname := symbol_graph_data_files[0]
	sg_numeric_adj := [...][]int32{
		[]int32{2, 7, 1},
		[]int32{4, 7, 0},
		[]int32{7, 0, 6, 5, 4, 3},
		[]int32{9, 6, 2},
		[]int32{1, 5, 7, 2},
		[]int32{4, 2, 6},
		[]int32{9, 8, 3, 2, 5},
		[]int32{1, 2, 4, 0},
		[]int32{9, 6},
		[]int32{6, 8, 3},
	}

	// load the symbol graph
	sg, err := LoadFromFile(fname, " ")
	if err != nil {
		t.Logf("error while creating symbol-graph from '%s'. reason: '%s'\n", fname, err)
		t.Fail()
	}

	// compare the symbol-graph's adjacency list
	for g, v := sg.G(), int32(0); v < g.V(); v++ {
		v_adj := g.Adj(v)
		ok := slice_utils.CmpInt32Slice(&sg_numeric_adj[v], &v_adj)

		if !ok {
			t.Logf("vertex: %d, got-adj: %v, want-adj: %v\n", v, sg_numeric_adj[v], g.Adj(v))
			t.Fail()
		}
	}
	return
}

//
// load a symbol graph from a file, and compare the adjacency list of
// each vertex in the graph with the expected adjacency
//
func TestNameAdj(t *testing.T) {
	fname := symbol_graph_data_files[0]
	var sg_name_adj = []struct {
		source string
		adj    []string
	}{
		{"JFK", []string{"MCO", "ATL", "ORD"}},
		{"MCO", []string{"ATL", "JFK", "HOU"}},
		{"ORD", []string{"ATL", "JFK", "PHX", "DFW", "HOU", "DEN"}},
		{"DEN", []string{"LAS", "PHX", "ORD"}},
		{"HOU", []string{"MCO", "DFW", "ATL", "ORD"}},
		{"DFW", []string{"HOU", "ORD", "PHX"}},
		{"PHX", []string{"LAS", "LAX", "DEN", "ORD", "DFW"}},
		{"ATL", []string{"MCO", "ORD", "HOU", "JFK"}},
		{"LAX", []string{"LAS", "PHX"}},
		{"LAS", []string{"PHX", "LAX", "DEN"}},
	}

	// load the symbol graph
	sg, err := LoadFromFile(fname, " ")
	if err != nil {
		t.Logf("error while creating symbol-graph from '%s'. reason: '%s'\n", fname, err)
		t.Fail()
	}

	// return the vertex name for a given vertex, or fail...
	get_vertex_name_or_fail := func(v int32) (name string) {
		var err error

		if name, err = sg.Name(v); err != nil {
			t.Logf("vertex: %d, error getting name. reason: %s\n", v, err)
			t.Fail()
		}
		return
	}

	// compare symbol-graph's named adjacency list
	for g, v := sg.G(), int32(0); v < g.V(); v++ {
		want := sg_name_adj[v]

		// source-name ok ?
		got_source := get_vertex_name_or_fail(v)
		if got_source != want.source {
			t.Logf("vertex-id: %d, got-name: %s, want-name: %s\n", v, got_source, want.source)
			t.Fail()
		}

		// create named adjacency list
		number_adj_list := g.Adj(v)
		named_adj := make([]string, len(number_adj_list))
		for i, v := range number_adj_list {
			named_adj[i] = get_vertex_name_or_fail(v)
		}

		// and compare
		ok := slice_utils.RelaxedCmpStringSlice(&want.adj, &named_adj)
		if !ok {
			t.Logf("vertex: %s(%d), got-adj: %v, want-adj: %v\n", v, named_adj, want.adj)
			t.Fail()
		}
	}

	return
}
