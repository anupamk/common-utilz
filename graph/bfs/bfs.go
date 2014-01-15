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
// provides an implementation of breadth-first-search for graphs
//
package bfs

import (
	"github.com/anupamk/common-utilz/graph"
	"github.com/anupamk/common-utilz/queue"
	"github.com/anupamk/common-utilz/stack"
)

//
// a BreadthFirstSearch object enumerating a set of vertices that can
// be visited from a given source vertex
//
type BreadthFirstSearch struct {
	visited []bool
	edge_to []int32
	source  int32
	count   int32
}

func New(G *graph.Graph, source int32) (bfs *BreadthFirstSearch) {
	bfs = &BreadthFirstSearch{
		visited: make([]bool, G.V()),
		edge_to: make([]int32, G.V()),
		source:  source,
		count:   0,
	}

	bfs.run_bfs(G, source)
	return
}

func (bfs *BreadthFirstSearch) Visited(v int32) bool { return bfs.visited[v] }
func (bfs *BreadthFirstSearch) Count() int32         { return bfs.count }

//
// returns the list of vertices that are connected to the source
// vertex
//
func (bfs *BreadthFirstSearch) ConnectedVertices() (vertex_list []int32) {
	vertex_list = make([]int32, bfs.count)
	for i, j := 0, 0; i < len(bfs.visited); i++ {
		if bfs.visited[i] == true {
			vertex_list[j] = int32(i)
			j = j + 1
		}
	}

	return
}

//
// returns true if a path source -> destination exists in the
// DepthFirstSearch object
//
func (bfs *BreadthFirstSearch) PathExistsTo(dest int32) bool {
	return bfs.visited[dest]
}

//
// enumerate the source -> destination path in the DepthFirstSearch
// object
//
func (bfs *BreadthFirstSearch) Path(dest int32) (path []int32) {
	if !bfs.PathExistsTo(dest) {
		return
	}

	// walk up the stack from current vertex to the source
	path_stack := stack.New()
	for w := dest; w != bfs.source; w = bfs.edge_to[w] {
		path_stack.Push(w)
	}

	path_stack.Push(bfs.source)

	// populate the path
	path = make([]int32, path_stack.Len())
	for i := 0; !path_stack.Empty(); i++ {
		val := path_stack.Pop()
		path[i] = val.(int32)
	}

	return
}

//
// private unexported stuff
//

// implements the canonical bfs procedure
func (bfs *BreadthFirstSearch) run_bfs(G *graph.Graph, source int32) {
	visited_queue := queue.New()

	bfs.visited[source] = true
	visited_queue.Push(source)

	for !visited_queue.Empty() {
		v := visited_queue.Pop().(int32)
		bfs.count++

		for _, w := range G.Adj(v) {
			if !bfs.visited[w] {
				bfs.edge_to[w] = v

				bfs.visited[w] = true
				visited_queue.Push(w)
			}
		}
	}

	return
}
