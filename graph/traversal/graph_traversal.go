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
	"fmt"
	"github.com/anupamk/common-utilz/graph"
	"github.com/anupamk/common-utilz/queue"
	"github.com/anupamk/common-utilz/stack"
)

type Walker struct {
	graph   *graph.Graph
	visited []bool
}

type Edge struct {
	src int32
	dst int32
}

// stringified representation of an edge v-w
func (edge Edge) String() string {
	str := ""
	str += fmt.Sprintf("%d-%d", edge.src, edge.dst)
	return str
}

var (
	EOG  = errors.New("End Of Graph")
	EOGS = errors.New("End Of Graph Subset")
)

type walker_style_t int

const (
	DO_DFS walker_style_t = iota
	DO_BFS                = iota
)

type GraphWalker func() (Edge, error)
type GraphSubsetWalker func() (Edge, error)

//
// this function returns a GraphSubsetWalker function. its repeated
// invokation traverses all edges in the subset-graph in a
// breadth-first-order.
//
// an 'EOGS' or end-of-graph-subset is returned, when all edges are
// visited.
//
func BFSGraphSubsetWalker(G *graph.Graph, source int32) GraphSubsetWalker {
	var ss_walker GraphSubsetWalker
	queue := queue.New()
	walker := &Walker{
		graph:   G,
		visited: make([]bool, G.V()),
	}

	visit_vertex := func(edge Edge) {
		walker.visited[edge.dst] = true
		queue.Push(edge)
		return
	}
	visit_vertex(Edge{source, source})

	ss_walker = func() (edge Edge, err error) {
		switch queue.Empty() {
		case false:
			// canonical bfs procedure
			edge = queue.Pop().(Edge)
			// fmt.Printf("popped-edge: %s\n", edge)

			for _, w := range walker.graph.Adj(edge.dst) {
				if !walker.visited[w] {
					visit_vertex(Edge{edge.dst, w})
				}
			}
		case true:
			err = EOGS
		}
		return
	}
	return ss_walker
}

//
// this function returns a GraphSubsetWalker function. its repeated
// invokation traverses all edges in the subset-graph in a
// depth-first-order.
//
// an 'EOGS' or end-of-graph-subset is returned, when all edges are
// visited.
//
func DFSGraphSubsetWalker(G *graph.Graph, source int32) GraphSubsetWalker {
	var ss_walker GraphSubsetWalker

	walker := &Walker{
		graph:   G,
		visited: make([]bool, G.V()),
	}

	// add source
	stack_1 := stack.New()
	stack_2 := stack.New()
	stack_1.Push(Edge{source, source})

	// complete the discovery
	for !stack_1.Empty() {
		edge := stack_1.Pop().(Edge)

		if !walker.visited[edge.dst] {
			walker.visited[edge.dst] = true
			stack_2.Push(edge)

			for _, w := range walker.graph.Adj(edge.dst) {
				e1 := Edge{edge.dst, w}
				stack_1.Push(e1)
			}
		}
	}

	// do the rest
	ss_walker = func() (next Edge, err error) {
		switch stack_2.Empty() {
		case false:
			next = stack_2.Pop().(Edge)

		case true:
			err = EOGS
		}

		return
	}

	return ss_walker
}

//
// this function returns a GraphWalker function type. repeated
// invokation of which traverses all edges of the graph in a
// breadth-first-order.
//
// an 'EOG' or end-of-graph is returned when all edges are visited.
//
func BFSGraphWalker(G *graph.Graph) GraphWalker {
	return create_fullgraph_walker(G, DO_BFS)
}

//
// this function returns a GraphWalker function type. repeated
// invokation of which traverses all vertices of the graph in
// depth-first-order.
//
// an 'EOG' or end-of-graph is returned when all edges are visited.
//
func DFSGraphWalker(G *graph.Graph) GraphWalker {
	return create_fullgraph_walker(G, DO_DFS)
}

// private un-exported stuff

//
// this function creates a graph-walker depending on the 'style' of
// walk (either bfs or dfs) to be done on the graph.
//
// just walk all the subsets of the graph using the appropriate subset
// walker.
//
func create_fullgraph_walker(G *graph.Graph, howto_walk walker_style_t) GraphWalker {
	var the_walker GraphWalker
	var num_visited_nodes int32
	var the_ss_walker GraphSubsetWalker

	visited_nodes := make([]bool, G.V())
	source := int32(0)

	switch howto_walk {
	case DO_BFS:
		the_ss_walker = BFSGraphSubsetWalker(G, source)

	case DO_DFS:
		the_ss_walker = DFSGraphSubsetWalker(G, source)
	}

	the_walker = func() (next Edge, err error) {
		next, err = the_ss_walker()
		visited_nodes[next.dst] = true
		num_visited_nodes += 1

		switch {
		case err == EOGS && num_visited_nodes < G.V():
			// find first available un-visited node.
			for v := int32(0); v < G.V(); v++ {
				if !visited_nodes[v] {
					the_ss_walker = BFSGraphSubsetWalker(G, v)

					// quench it...
					_, _ = the_ss_walker()
					break
				}
			}

		case err == EOGS && num_visited_nodes >= G.V():
			// we are done here
			err = EOG
		}
		return
	}

	return the_walker
}
