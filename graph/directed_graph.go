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
// provides adjacency list based implementation of directed graphs or
// digraphs.
//
package graph

import (
	"container/list"
)

//
// adjacency list representation of a Graph, which contains 'V'
// vertices and 'E' edges. vertices are in the range {0, V-1} for ease
// of processing
//
type Digraph struct {
	v   int32
	e   int32
	adj []list.List
}

//
// this function is called to create a new skeleton graph, with a
// specific number of vertices
//
func CreateDigraph(V int32) *Digraph {
	return &Digraph{
		v:   V,
		e:   0,
		adj: make([]list.List, V),
	}
}

func (G *Digraph) V() int32 { return G.v }
func (G *Digraph) E() int32 { return G.e }

//
// return the list of vertices adjacent to a given vertex 'v'.
//
func (G *Digraph) Adj(v int32) []int32 {
	adj_list := G.adj[v]
	vertex_list := make([]int32, adj_list.Len())

	for i, node := 0, adj_list.Front(); node != nil; i, node = i+1, node.Next() {
		v := node.Value.(int32)
		vertex_list[i] = v
	}

	return vertex_list
}

//
// in a digraph G, add an edge between vertices 'v' and 'w'.
//
func (G *Digraph) AddEdge(v, w int32) {
	G.adj[v].PushFront(w)
	G.e += 1

	return
}

//
// return the reverse of a digraph i.e. adjacency list of each vertex
// is reversed
//
func (G *Digraph) Reverse() (RevG *Digraph) {
	RevG = CreateDigraph(G.V())
	for v := int32(0); v < G.V(); v++ {
		for _, w := range G.Adj(v) {
			RevG.AddEdge(w, v)
		}
	}

	return
}

func (G *Digraph) String() string { return graph_stringifier(G) }

//
// enumerate some fundamental properties of a graph
//
func (G *Digraph) Degree(v int32) (degree int32) { return int32(len(G.Adj(v))) }
func (G *Digraph) AverageDegree() float64        { return float64(2 * G.V() / G.E()) }

func (G *Digraph) MaxDegree() (max_degree int32) {
	max_degree = 0

	for v := int32(0); v < G.V(); v++ {
		d := G.Degree(v)
		if d > max_degree {
			max_degree = d
		}
	}

	return
}
