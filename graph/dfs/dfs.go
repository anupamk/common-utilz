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
// provides an implementation of depth-first-search algorithm for
// graphs
//
package dfs

import (
	"github.com/anupamk/common-utilz/graph"
)

type DepthFirstSearch struct {
	visited []bool
	count   int32
}

func New(G *graph.Graph, source int32) (dfs *DepthFirstSearch) {
	dfs = &DepthFirstSearch{
		visited: make([]bool, G.V()),
		count:   0,
	}

	dfs.run_dfs(G, source)
	return
}

func (dfs *DepthFirstSearch) Visited(v int32) bool { return dfs.visited[v] }
func (dfs *DepthFirstSearch) Count() int32         { return dfs.count }

//
// returns the list of vertices that are connected to the source
// vertex
//
func (dfs *DepthFirstSearch) ConnectedVertices() (vertex_list []int32) {
	vertex_list = make([]int32, dfs.count)
	for i, j := 0, 0; i < len(dfs.visited); i++ {
		if dfs.visited[i] == true {
			vertex_list[j] = int32(i)
			j = j + 1
		}
	}

	return
}

//
// private unexported stuff
//

// implements the canonical recursive dfs procedure
func (dfs *DepthFirstSearch) run_dfs(G *graph.Graph, source int32) {
	dfs.visited[source] = true
	dfs.count += 1

	for _, w := range G.Adj(source) {
		if dfs.visited[w] == false {
			dfs.run_dfs(G, w)
		}
	}

	return
}
