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
package traversal

import (
	"errors"
	"github.com/anupamk/common-utilz/graph"
	"github.com/anupamk/common-utilz/queue"
)

type Walker struct {
	graph   *graph.Graph
	visited []bool
}
type GraphWalker func() (int32, error)

var (
	EOG = errors.New("End Of Graph") // cheeky
)

//
// this function is called returns GraphWalker which walks a given
// graph 'G' starting from 'source' vertex in a
// breadth-first-order.
//
// each invokation of GraphWalker returns the visited-node. when all
// nodes are visited, an 'EOG' or end-of-graph is returned.
//
func BFSWalker(G *graph.Graph, source int32) GraphWalker {
	walker := &Walker{
		graph:   G,
		visited: make([]bool, G.V()),
	}
	var bfs_walker GraphWalker
	var visited_nodes int32
	queue := queue.New()

	// visit a vertex
	visit_vertex := func(source int32) {
		walker.visited[source] = true
		visited_nodes += 1
		queue.Push(source)
		return
	}

	visit_vertex(source)

	bfs_walker = func() (int32, error) {
		var next_node int32
		var err error

	restart_walk:
		switch queue.Empty() {
		case false:
			// the canonical bfs procedure
			next_node = queue.Pop().(int32)
			for _, w := range walker.graph.Adj(next_node) {
				if !walker.visited[w] {
					visit_vertex(w)
				}
			}

		case true && visited_nodes < walker.graph.V():
			// enqueue first node that is not processed yet...
			for v := int32(0); v < walker.graph.V(); v++ {
				if !walker.visited[v] {
					visit_vertex(v)
					goto restart_walk
				}
			}

		case true && visited_nodes == walker.graph.V():
			err = EOG
		}

		return next_node, err
	}

	return bfs_walker
}

//
// this function returns a GraphWalker, which traverses the graph 'G'
// from source-vertex 'source' in depth-first-order.
//
// each invokation of GraphWalker returns the visited-node. when all
// nodes are visited, an 'EOG' or end-of-graph is returned.
//
func DFSWalker(G *graph.Graph, source int32) GraphWalker {
	walker := &Walker{
		graph:   G,
		visited: make([]bool, G.V()),
	}
	var dfs_walker GraphWalker
	var visited_nodes int32

	// the dfs 'stack' depth is 1. this is because, for every
	// vertext traversed in the dfs order, the walker relinquishes
	// control to the caller. so, nothing fancy is required
	// here, just a little bit of bookeeping
	var dfs_stack_empty bool
	var dfs_stack int32

	// visit a vertex
	visit_vertex := func(v int32) {
		walker.visited[v] = true
		visited_nodes += 1
		dfs_stack = v
		dfs_stack_empty = false
		return
	}
	visit_vertex(source)

	dfs_walker = func() (int32, error) {
		var next_node int32
		var err error

	restart_dfs_walk:
		switch dfs_stack_empty {
		case false:
			// canonical dfs procedure
			next_node = dfs_stack
			for _, w := range walker.graph.Adj(next_node) {
				if !walker.visited[w] {
					visit_vertex(w)
					return next_node, err
				}
			}

			// since we are here, it implies that all
			// reachable vertices (from original source)
			// have been visited. mark the stack as
			// such...
			dfs_stack_empty = true

		case true && visited_nodes < walker.graph.V():
			// push first unvisited node, and process it.
			for v := int32(0); v < walker.graph.V(); v++ {
				if !walker.visited[v] {
					visit_vertex(v)
					goto restart_dfs_walk
				}
			}

		case true && visited_nodes == walker.graph.V():
			err = EOG
		}

		return next_node, err
	}

	return dfs_walker
}
