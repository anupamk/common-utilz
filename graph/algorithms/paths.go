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
package algorithms

import (
	"fmt"
	"github.com/anupamk/common-utilz/graph"
	"github.com/anupamk/common-utilz/graph/traversal"
	"github.com/anupamk/common-utilz/stack"
)

type edge_to_t struct {
	v     int32
	valid bool
}

// stringified representation of an edge_to_t instance
func (e edge_to_t) String() string {
	str := ""
	if e.valid {
		str += fmt.Sprintf("%d", e.v)
	} else {
		str += fmt.Sprintf("%d: not-connected", e.v)
	}

	return str
}

type GraphPath struct {
	edge_to []edge_to_t
	source  int32
}

// stringified representation of graph path
func (path *GraphPath) String() string {
	str := ""
	str += fmt.Sprintf("source-vertex: %d\n", path.source)
	for i, e := range path.edge_to {
		str += fmt.Sprintf("%d: %s\n", i, e)
	}

	return str
}

//
// this function is called to compute the depth-first-search path for
// the given graph from a source vertex
//
func DFSPath(G *graph.Graph, source int32) (path *GraphPath) {
	path = &GraphPath{
		edge_to: make([]edge_to_t, G.V()),
		source:  source,
	}
	ss_walker := traversal.DFSGraphSubsetWalker(G, source)

	for edge, err := ss_walker(); err != traversal.EOGS; edge, err = ss_walker() {
		path.edge_to[edge.Dst] = edge_to_t{edge.Src, true}
	}

	return
}

//
// this function is called to compute the breadth-first-search path
// for the given graph from a source vertex
//
func BFSPath(G *graph.Graph, source int32) (path *GraphPath) {
	path = &GraphPath{
		edge_to: make([]edge_to_t, G.V()),
		source:  source,
	}
	ss_walker := traversal.BFSGraphSubsetWalker(G, source)

	for edge, err := ss_walker(); err != traversal.EOGS; edge, err = ss_walker() {
		path.edge_to[edge.Dst] = edge_to_t{edge.Src, true}
	}

	return
}

// some commonly used queries on paths

//
// this function returns true if a path from source -> dst exists,
// false otherwise. signals an error if the dst doesn't seem to be
// valid.
//
func (gp *GraphPath) HasPathTo(dst int32) (yesno bool, err error) {
	if dst > int32(len(gp.edge_to)) {
		err = fmt.Errorf("bogus destination: %d\n", dst)
	}

	yesno = gp.edge_to[dst].valid
	return
}

//
// this function enumerates the path from source -> dst if such a path
// exists.
//
func (gp *GraphPath) PathTo(dst int32) (path []int32, err error) {
	var path_exists bool

	// invalid destination
	if path_exists, err = gp.HasPathTo(dst); err != nil {
		return
	}

	// no paths exist
	if !path_exists {
		err = fmt.Errorf("no path to: %d\n", dst)
		return
	}

	// find the path
	stack := stack.New()
	for v := dst; v != gp.source; v = gp.edge_to[v].v {
		stack.Push(v)
	}
	stack.Push(gp.source)

	path = make([]int32, stack.Len())
	for i := 0; i < len(path); i++ {
		path[i] = stack.Pop().(int32)
	}

	return
}
