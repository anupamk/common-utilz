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
	"github.com/anupamk/common-utilz/stack"
)

type Walker struct {
	graph   *graph.Graph
	visited []bool
}

type GraphWalker func() (int32, error)
type GraphSubsetWalker func() (int32, error)

var (
	EOG = errors.New("End Of Graph") // cheeky
	EOS = errors.New("End Of SubSet")
)

//
// this function returns a GraphWalker function type. repeated
// invokation of GraphWalker walks all vertices of the graph in
// breadth-first-order.
//
// an 'EOG' or end-of-graph is returned to the caller, when all
// vertices have been visited...
//
func BFSGraphWalker(G *graph.Graph) GraphWalker {
	var bfs_walker GraphWalker
	var visited_nodes int32 // number of nodes visited thus far

	walker := &Walker{
		graph:   G,
		visited: make([]bool, G.V()),
	}
	queue := queue.New()

	// since we traverse all the nodes of the graph, choice of
	// initial source vertex doesn't matter, so we pick one which
	// is always guarenteed to be there.
	source := int32(0)

	// visit a vertex
	visit_vertex := func(source int32) {
		walker.visited[source] = true
		visited_nodes += 1
		queue.Push(source)
		return
	}
	visit_vertex(source)

	bfs_walker = func() (next_node int32, err error) {

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

		return
	}

	return bfs_walker
}

//
// this function returns a GraphSubsetWalker function. its repeated
// invokation visits all reachable vertices in the graph (from source
// vertex) in breadth-first-order.
//
// fwiw, this is a slightly simpler version of the BFSGraphWalker
// function above.
//
// an 'EOS' or end-of-graph-subset is returned, when all vertices are
// visited...
//
func BFSGraphSubsetWalker(G *graph.Graph, source int32) GraphSubsetWalker {
	var bfs_ss_walker GraphSubsetWalker

	walker := &Walker{
		graph:   G,
		visited: make([]bool, G.V()),
	}
	queue := queue.New()

	// visit a vertex
	visit_vertex := func(v int32) {
		walker.visited[v] = true
		queue.Push(v)
		return
	}
	visit_vertex(source)

	bfs_ss_walker = func() (next_node int32, err error) {
		switch queue.Empty() {
		case false:
			// the canonical bfs procedure
			next_node = queue.Pop().(int32)
			for _, w := range walker.graph.Adj(next_node) {
				if !walker.visited[w] {
					visit_vertex(w)
				}
			}

		case true:
			err = EOS
		}

		return
	}

	return bfs_ss_walker
}

//
// this function returns a GraphWalker function type. repeated
// invokation of GraphWalker walks all vertices of the graph in
// depth-first-order.
//
// an 'EOG' or end-of-graph is returned to the caller, when all
// vertices have been visited...
//
func DFSGraphWalker(G *graph.Graph) GraphWalker {
	var dfs_walker GraphWalker
	var visited_nodes int32

	walker := &Walker{
		graph:   G,
		visited: make([]bool, G.V()),
	}
	stack := stack.New()

	// since we traverse all the nodes of the graph, choice of
	// initial source vertex doesn't matter, so we pick one which
	// is always guarenteed to be there.
	source := int32(0)

	// visit a vertex
	visit_vertex := func(v int32) {
		walker.visited[v] = true
		visited_nodes += 1
		stack.Push(v)
		return
	}
	visit_vertex(source)

	dfs_walker = func() (next_node int32, err error) {
	restart_dfs_walk:
		switch stack.Empty() {
		case false:
			// canonical dfs procedure
			next_node = stack.Pop().(int32)
			for _, w := range walker.graph.Adj(next_node) {
				if !walker.visited[w] {
					visit_vertex(w)
				}
			}

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

//
// this function returns a GraphSubsetWalker function. its repeated
// invokation visits all reachable vertices in the graph (from source
// vertex) in depth-first-order.
//
// fwiw, this is a slightly simpler version of the DFSGraphWalker
// function above.
//
// an 'EOS' or end-of-graph-subset is returned, when all vertices are
// visited...
//
func DFSGraphSubsetWalker(G *graph.Graph, source int32) GraphSubsetWalker {
	var dfs_ss_walker GraphSubsetWalker

	walker := &Walker{
		graph:   G,
		visited: make([]bool, G.V()),
	}
	stack := stack.New()

	// visit a vertex
	visit_vertex := func(v int32) {
		walker.visited[v] = true
		stack.Push(v)
		return
	}
	visit_vertex(source)

	dfs_ss_walker = func() (next_node int32, err error) {
		switch stack.Empty() {
		case false:
			// canonical dfs procedure
			next_node = stack.Pop().(int32)
			for _, w := range walker.graph.Adj(next_node) {
				if !walker.visited[w] {
					visit_vertex(w)
				}
			}

		case true:
			err = EOS
		}

		return
	}

	return dfs_ss_walker
}
