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
)

//
// a BreadthFirstSearch object enumerating a set of vertices that can
// be visited from a given source vertex
//
type BreadthFirstSearch struct {
	visited []bool
	count   int32
}

func New(G *graph.Graph, source int32) (bfs *BreadthFirstSearch) {
	bfs = &BreadthFirstSearch{
		visited: make([]bool, G.V()),
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
// private unexported stuff
//

// implements the canonical bfs procedure
func (bfs *BreadthFirstSearch) run_bfs(G *graph.Graph, source int32) {
	visited_queue := queue.New()
	visited_queue.Push(source)

	for !visited_queue.Empty() {
		v := visited_queue.Pop().(int32)

		// mark current vertex as visited, and add it's
		// adjacent vertices to the queue
		if bfs.visited[v] == false {
			bfs.visited[v] = true
			bfs.count += 1

			for _, w := range G.Adj(v) {
				visited_queue.Push(w)
			}
		}
	}

	return
}
