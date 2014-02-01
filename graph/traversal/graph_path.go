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
// in this package we try to separate out graph traversal-order from
// actual procedures which use these traversals.
//
// this file implements the shortest-path for undirected graph using
// the traversal api's
//
package traversal

import (
	"fmt"
	"github.com/anupamk/common-utilz/graph"
	"github.com/anupamk/common-utilz/stack"
)

type valid_edge_to_t struct {
	from int32
	ok   bool
}

func (edge valid_edge_to_t) String() (str string) {
	if !edge.ok {
		return
	}

	str += fmt.Sprintf("%d\n", edge.from)
	return
}

// this structure represents paths to all vertices in the graph from a
// given source
type SingleSourcePaths struct {
	source  int32
	edge_to []valid_edge_to_t
}

// determine shortest paths from source to all other vertices in the
// graph.
func SingleSourceShortestPaths(g *graph.Graph, source int32) (ssp *SingleSourcePaths, sp_err error) {
	if source > g.V() {
		sp_err = fmt.Errorf("bad source-vertex: '%d'\n", source)
		return
	}

	ssp = &SingleSourcePaths{
		source:  source,
		edge_to: make([]valid_edge_to_t, g.V()),
	}

	// just run bfs from the source for the recently discovered
	// edge, destination is reachable from the source...
	bfs_walker := BFSGraphSubsetWalker(g, source)
	for E, err := bfs_walker(); err != EOGS; E, err = bfs_walker() {
		ssp.edge_to[E.Dst].from = E.Src
		ssp.edge_to[E.Dst].ok = true
	}

	return
}

//
// this function returns true if a path-to a given destination node
// exists, false otherwise.
//
func (ssp *SingleSourcePaths) PathExists(dest int32) (yesno bool) {
	return ssp.edge_to[dest].ok
}

//
// this function enumerates the path to a given destination from the
// source node. panics if no path exists...
//
func (ssp *SingleSourcePaths) PathTo(dest int32) []int32 {
	if !ssp.PathExists(dest) {
		err := fmt.Errorf("no path to '%d' exists\n", dest)
		panic(err)
	}

	// determine the path (dest -> source)
	path_stack := stack.New()
	for v := dest; v != ssp.source; v = ssp.edge_to[v].from {
		// should never happen
		if !ssp.edge_to[v].ok {
			err := fmt.Errorf("vertex %d, path exists, but is invalid !!!", v)
			panic(err)
		}
		path_stack.Push(v)
	}
	path_stack.Push(ssp.source)

	// create source -> dest result
	path := make([]int32, path_stack.Len())
	for i := 0; !path_stack.Empty(); i++ {
		path[i] = path_stack.Pop().(int32)
	}

	return path
}
